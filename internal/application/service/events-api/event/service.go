package event

import (
	"github.com/Vovchikus/events-api/internal/domain/event/model"
	"github.com/Vovchikus/events-api/internal/infrastructure/db/kafka/event"
	"github.com/Vovchikus/events-api/internal/infrastructure/logger"
)

type (
	Service interface {
		CreateEvent(event *model.Event)
	}

	service struct {
		kafkaEventRepository event.Repository
		logger               logger.Interface
	}
)

func NewEventAppService(r event.Repository, l logger.Interface) *service {
	return &service{r, l}
}

func (s *service) CreateEvent(event *model.Event) {
	_ = s.kafkaEventRepository.CreateEvent(event)
}
