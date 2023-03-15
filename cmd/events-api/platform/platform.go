package platform

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/Vovchikus/events-api/internal/application/service/events-api/event"
	"github.com/Vovchikus/events-api/internal/infrastructure/config"
	"github.com/Vovchikus/events-api/internal/infrastructure/logger"
	eventsApiApplication "github.com/Vovchikus/events-api/internal/presentation/events-api/v1/event"
	"google.golang.org/grpc"
	"os"

	eventsRepo "github.com/Vovchikus/events-api/internal/infrastructure/db/kafka/event"
	eventsApi "github.com/Vovchikus/events-api/pkg/events-api"
)

var (
	ConfigPath = flag.String("config-path", ".", "Config path")
	ConfigType = flag.String("config-type", "env", "Config type")
)

type App struct {
	Server       *grpc.Server
	Logger       logger.Interface
	Config       config.ViperConfig
	EventService event.Service
}

func New() *App {

	c := config.ViperConfig{ConfigPath: *ConfigPath, ConfigType: *ConfigType}
	_ = c.Init()

	loggerConfig := logger.Configuration{
		EnableConsole:      c.GetConfig().LoggerConsole.Enable,
		ConsoleLevel:       c.GetConfig().LoggerConsole.Level,
		ConsoleJSONFormat:  c.GetConfig().LoggerConsole.Json,
		EnableFile:         c.GetConfig().LoggerFile.Enable,
		FileLevel:          c.GetConfig().LoggerFile.Level,
		FileJSONFormat:     c.GetConfig().LoggerFile.Json,
		FileLocation:       c.GetConfig().LoggerFile.Path,
		FileLocationPrefix: "grpc",
	}
	l := logger.NewZapLogger(loggerConfig)

	defer func() {
		if r := recover(); r != nil {
			l.WithFields(logger.Fields{"error": r}).Panic("Program fatal error")
			os.Exit(1)
		}
	}()

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true

	host := c.GetConfig().Kafka.Host
	fmt.Println(host)

	kafkaClient, err := sarama.NewClient([]string{c.GetConfig().Kafka.Host + ":" + c.GetConfig().Kafka.Port}, kafkaConfig)
	if err != nil {
		fmt.Println("Ошибка при создании клиента:", err.Error())
	}

	producer, err := sarama.NewSyncProducerFromClient(kafkaClient)
	if err != nil {
		fmt.Println("Ошибка при создании производителя:", err.Error())
	}

	s := grpc.NewServer()
	eventsApi.RegisterEventServiceServer(
		s,
		eventsApiApplication.NewEventApiService(
			event.NewEventAppService(
				eventsRepo.NewKafkaEventRepository(producer, l),
				l,
			),
		),
	)

	return &App{
		Server: s,
		Logger: l,
		Config: c,
		EventService: event.NewEventAppService(
			eventsRepo.NewKafkaEventRepository(producer, l),
			l,
		),
	}
}
