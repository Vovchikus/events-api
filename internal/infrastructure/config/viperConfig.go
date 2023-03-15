package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ViperConfigInterface interface {
	Init() error
	GetString(key string, def ...string) string
	GetBool(key string) bool
	GetInt(key string, def ...int) int
	GetFloat64(key string, def ...float64) float64
	GetConfig() BaseConfig
}

type ViperConfig struct {
	ConfigPath string
	ConfigType string `default:"env"`
	BaseConfig BaseConfig
}

func (config *ViperConfig) Init() error {
	viper.AddConfigPath(config.ConfigPath)
	viper.SetConfigType(config.ConfigType)
	var err error
	if config.isEnv() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()
		err = viper.ReadInConfig()
	} else {
		err = viper.ReadInConfig()
	}
	config.initConfig()

	return err
}

func (config *ViperConfig) initConfig() {

	config.BaseConfig = BaseConfig{
		LoggerConsole: getLoggerConsoleConfig(),
		LoggerFile:    getLoggerFileConfig(),
		Grpc:          getGrpcConfig(),
		Kafka:         getKafkaConfig(),
	}
}

func (config *ViperConfig) GetString(key string, def ...string) string {
	value := viper.GetString(config.prepareKey(key))
	if value == "" && len(def) > 0 {
		value = def[0]
	}

	return value
}

func (config *ViperConfig) GetBool(key string) bool {
	return viper.GetBool(config.prepareKey(key))
}

func (config *ViperConfig) GetInt(key string, def ...int) int {
	value := viper.GetInt(config.prepareKey(key))
	if value == 0 && len(def) > 0 {
		value = def[0]
	}

	return value
}

func (config *ViperConfig) GetFloat64(key string, def ...float64) float64 {
	value := viper.GetFloat64(config.prepareKey(key))
	if value == 0 && len(def) > 0 {
		value = def[0]
	}

	return value
}

func (config *ViperConfig) GetConfig() BaseConfig {
	return config.BaseConfig
}

func (config *ViperConfig) prepareKey(key string) string {
	if config.isEnv() {
		key = strings.ReplaceAll(key, ".", "_")
		key = strings.ToUpper(key)
	}

	return key
}

func (config *ViperConfig) isEnv() bool {
	return config.ConfigType == "env"
}

func isEnvTrue(s string) bool {
	return os.Getenv(s) == "true"
}

func isEnvEmpty(s string) bool {
	return os.Getenv(s) == ""
}

func getGrpcConfig() GrpcConfig {
	var grpcGatewayEnable, grpcSaveRequest, grpcSaveResponse bool
	if isEnvTrue("GRPC_GATEWAY_ENABLE") {
		grpcGatewayEnable = true
	}
	if isEnvTrue("GRPC_SAVE_REQUEST") {
		grpcSaveRequest = true
	}
	if isEnvTrue("GRPC_SAVE_RESPONSE") {
		grpcSaveResponse = true
	}
	return GrpcConfig{
		Network: os.Getenv("GRPC_NETWORK"),
		Address: os.Getenv("GRPC_ADDRESS"),
		Port:    os.Getenv("GRPC_PORT"),
		Gateway: GrpcGatewayConfig{
			Enable:  grpcGatewayEnable,
			Address: os.Getenv("GRPC_GATEWAY_ADDRESS"),
			Port:    os.Getenv("GRPC_GATEWAY_PORT"),
		},
		SaveRequest:  grpcSaveRequest,
		SaveResponse: grpcSaveResponse,
	}
}

func getKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Port: os.Getenv("KAFKA_PORT"),
		Host: os.Getenv("KAFKA_HOST"),
	}
}

func getLoggerConsoleConfig() LoggerConfig {
	var loggerConsoleEnable, loggerConsoleJson bool
	if isEnvTrue("LOGGER_CONSOLE_ENABLE") {
		loggerConsoleEnable = true
	}
	if isEnvTrue("LOGGER_CONSOLE_JSON") {
		loggerConsoleJson = true
	}
	return LoggerConfig{
		Enable: loggerConsoleEnable,
		Level:  os.Getenv("LOGGER_CONSOLE_LEVEL"),
		Json:   loggerConsoleJson,
		Path:   "",
	}
}

func getLoggerFileConfig() LoggerConfig {
	var loggerFileEnable, loggerFileJson bool
	if isEnvTrue("LOGGER_FILE_ENABLE") {
		loggerFileEnable = true
	}
	if isEnvTrue("LOGGER_FILE_JSON") {
		loggerFileJson = true
	}
	return LoggerConfig{
		Enable: loggerFileEnable,
		Level:  os.Getenv("LOGGER_FILE_LEVEL"),
		Json:   loggerFileJson,
		Path:   os.Getenv("LOGGER_FILE_PATH"),
	}
}
