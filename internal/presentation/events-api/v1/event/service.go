package event

import (
	"github.com/Vovchikus/events-api/internal/application/service/events-api/event"
	pkg "github.com/Vovchikus/events-api/pkg/events-api"
	"google.golang.org/grpc"
)

type Implementation struct {
	pkg.UnimplementedEventServiceServer
	Server       *grpc.Server
	eventService event.Service
}

func NewEventApiService(e event.Service) *Implementation {
	return &Implementation{eventService: e}
}
