package platform

import (
	"encoding/json"
	eventsApi "github.com/Vovchikus/events-api/pkg/events-api"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go a.grpcService(&wg)
	if a.Config.GetConfig().Grpc.Gateway.Enable {
		wg.Add(1)
		go a.proxyService(&wg)
	}
	wg.Wait()
	return nil
}

func (a *App) grpcService(wg *sync.WaitGroup) {
	defer wg.Done()

	c := make(chan os.Signal, 1)

	go func() {
		oscall := <-c
		a.Logger.Info("system call:%+v", oscall)
		a.Server.GracefulStop()
		os.Exit(0)
	}()

	if err := a.Serve(); err != nil {
		a.Logger.Fatal("failed to serve", err)
	}
}

func (a *App) Serve() error {
	address := a.Config.GetConfig().Grpc.Address + ":" + a.Config.GetConfig().Grpc.Port
	listener, err := net.Listen(a.Config.GetConfig().Grpc.Network, address)
	if err != nil {
		return err
	}

	return a.Server.Serve(listener)
}

func (a *App) proxyService(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	customMarshaller := &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}
	muxOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, customMarshaller)
	mux := runtime.NewServeMux(muxOpt)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcAddress := ":" + a.Config.GetConfig().Grpc.Port
	a.Logger.Info("Start GRPC server on: " + a.Config.GetConfig().Grpc.Address + grpcAddress)

	err := eventsApi.RegisterEventServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		a.Logger.Fatal("error reg address endpoint", err)
	}

	gatewayAddress := ":" + a.Config.GetConfig().Grpc.Gateway.Port
	a.Logger.Info("Start Http server on: " + a.Config.GetConfig().Grpc.Gateway.Address + gatewayAddress)

	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Ready probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Ready probe is positive.")
	}()

	httpMux := http.NewServeMux()
	muxWithIP := loggingMiddleware(mux)
	httpMux.Handle("/", muxWithIP)
	httpMux.Handle("/live", http.HandlerFunc(liveHandle))
	httpMux.Handle("/health", http.HandlerFunc(healthHandle))
	a.Logger.Info("Liveness started")
	httpMux.Handle("/ready", readyHandle(isReady))
	a.Logger.Info("Readiness started")

	srv := &http.Server{
		Addr:    gatewayAddress,
		Handler: httpMux,
	}

	a.Logger.Fatal(srv.ListenAndServe().Error())
}

func liveHandle(w http.ResponseWriter, _ *http.Request) {
	body, err := json.Marshal("This is service live!")
	if err != nil {
		log.Fatalf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Fatalf("Could not write body: %v", err)
		return
	}
}

func healthHandle(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readyHandle(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ipAddress := req.RemoteAddr
		fwdAddress := req.Header.Get("X-Forwarded-For")
		if fwdAddress != "" {
			ipAddress = fwdAddress

			ips := strings.Split(fwdAddress, ", ")
			if len(ips) > 1 {
				ipAddress = ips[0]
			}
		}
		rw.Header().Set("X-Forwarded-For", ipAddress)
		h.ServeHTTP(rw, req)
	})
}
