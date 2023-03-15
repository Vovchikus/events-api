package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

// zapLogger - объект-обёртка над логгером Zap
type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// NewZapLogger - инициализация логгера
func NewZapLogger(config Configuration) Interface {
	var cores []zapcore.Core

	// Создает обработчик вывода логов в консоль
	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	// Создает обработчик сохранения логов в файл
	if config.EnableFile {
		filePath := filepath.Dir(config.FileLocation) + "/" + config.FileLocationPrefix + "/log-" + time.Now().Format("2006-01-02") + ".log"
		level := getZapLevel(config.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: filePath,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}
}

// Debug - вывести отладку
func (z *zapLogger) Debug(format string, args ...interface{}) {
	z.sugaredLogger.Debugf(format, args...)
}

// Info - вывести информацию
func (z *zapLogger) Info(format string, args ...interface{}) {
	z.sugaredLogger.Infof(format, args...)
}

// Warn - вывести предупреждение
func (z *zapLogger) Warn(format string, args ...interface{}) {
	z.sugaredLogger.Warnf(format, args...)
}

// Error - вывести ошибку
func (z *zapLogger) Error(format string, args ...interface{}) {
	z.sugaredLogger.Errorf(format, args...)
}

// Fatal - вывести критическую ошибку
func (z *zapLogger) Fatal(format string, args ...interface{}) {
	z.sugaredLogger.Fatalf(format, args...)
}

// Panic - вывести критическую ошибку
func (z *zapLogger) Panic(format string, args ...interface{}) {
	z.sugaredLogger.Fatalf(format, args...)
}

// Printf - вывести текст
func (z *zapLogger) Printf(format string, args ...interface{}) {
	z.sugaredLogger.Infof(format, args...)
}

// WithFields - вывести дополнительные данные
func (z *zapLogger) WithFields(fields Fields) Interface {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := z.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

// zapEncoder - установить формат даты
func zapEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05Z07:00"))
}

// getEncoder - установить правила форматирование логов
func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getZapLevel - конвертирует уровень логов в формат Zap
func getZapLevel(level string) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
