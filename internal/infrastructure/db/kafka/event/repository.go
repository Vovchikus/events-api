package event

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/Vovchikus/events-api/internal/domain/event/model"
	"github.com/Vovchikus/events-api/internal/infrastructure/logger"
)

type Repository interface {
	CreateEvent(event *model.Event) error
}

type kafkaEventRepository struct {
	producer sarama.SyncProducer
	logger   logger.Interface
}

func NewKafkaEventRepository(p sarama.SyncProducer, l logger.Interface) Repository {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	repo := &kafkaEventRepository{
		producer: p,
		logger:   l,
	}

	return repo
}

func (r *kafkaEventRepository) CreateEvent(event *model.Event) error {
	messageBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error marshalling event %v", err)
	}

	message := &sarama.ProducerMessage{
		Topic: "event",
		Value: sarama.ByteEncoder(messageBytes),
	}

	fmt.Println(message)

	_, _, err = r.producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("error sending event message: %v", err)
	}

	return nil
}
