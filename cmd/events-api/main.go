package main

import (
	"flag"
	"github.com/Vovchikus/events-api/cmd/events-api/platform"
	"github.com/Vovchikus/events-api/internal/infrastructure/logger"
	"runtime"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := platform.New()
	if err := app.Run(); err != nil {
		app.Logger.WithFields(logger.Fields{"error": err}).Fatal("Program fatal error")
	}
}
