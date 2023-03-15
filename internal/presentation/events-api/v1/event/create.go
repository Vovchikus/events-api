package event

import (
	"context"
	"github.com/Vovchikus/events-api/internal/domain/event/model"
	desc "github.com/Vovchikus/events-api/pkg/events-api"
)

func (i *Implementation) PostCreateEvent(ctx context.Context, request *desc.CreateEventRequest) (r *desc.CreateEventResponse, e error) {

	event := model.NewEvent(request.Id, request.Content)

	i.eventService.CreateEvent(event)

	response := desc.CreateEventResponse{}
	response.Content = event.Content
	response.Id = event.ID

	return &response, nil
}
