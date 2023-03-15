package config

type LoggerConfig struct {
	Enable bool
	Level  string
	Json   bool
	Path   string
}

type GrpcConfig struct {
	Network      string
	Address      string
	Port         string
	SaveRequest  bool
	SaveResponse bool
	Gateway      GrpcGatewayConfig
}

type GrpcGatewayConfig struct {
	Enable  bool
	Address string
	Port    string
}

type WorkersConfig struct {
	Houses    int
	Addresses int
}

type KafkaConfig struct {
	Host string
	Port string
}

type BaseConfig struct {
	LoggerConsole LoggerConfig
	LoggerFile    LoggerConfig
	Grpc          GrpcConfig
	Workers       WorkersConfig
	Kafka         KafkaConfig
}
