package event

import (
	"context"
	"fmt"
	"github.com/Vovchikus/events-api/internal/domain/event/model"
	desc "github.com/Vovchikus/events-api/pkg/events-api"
)

func (i *Implementation) GetEvents(ctx context.Context, request *desc.CreateEventRequest) (desc.CreateEventResponse, error) {

	fmt.Println(request)

	event := model.NewEvent("123", "content")

	i.eventService.CreateEvent(event)

	return desc.CreateEventResponse{}, nil
}
