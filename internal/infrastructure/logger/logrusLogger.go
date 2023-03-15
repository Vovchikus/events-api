package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Объект-обёртка над логгером Logrus
type logrusLogger struct {
	logger *logrus.Logger
}

// Объект-обёртка над логгером Logrus для сохранения логов с доп. полями
type logrusLogEntry struct {
	entry *logrus.Entry
}

// getFormatter - установить правила форматирование логов
func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}

	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

// NewLogrusLogger - Инициализация логгера
func NewLogrusLogger(config Configuration) Interface {
	logLevel := config.ConsoleLevel
	if logLevel == "" {
		logLevel = config.FileLevel
	}

	// Устанавливает уровень логирования
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	stdOutHandler := os.Stdout
	// Создает обработчик сохранения логов в файл
	filePath := filepath.Dir(config.FileLocation) + "/" + config.FileLocationPrefix + "/log-" + time.Now().Format("2006-01-02") + ".log"
	fileHandler := &lumberjack.Logger{
		Filename: filePath,
		MaxSize:  100,
		Compress: true,
		MaxAge:   28,
	}
	// Создает обработчик вывода логов в консоль
	lLogger := &logrus.Logger{
		Out:       stdOutHandler,
		Formatter: getFormatter(config.ConsoleJSONFormat),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	if config.EnableConsole && config.EnableFile {
		lLogger.SetOutput(io.MultiWriter(stdOutHandler, fileHandler))
	} else {
		if config.EnableFile {
			lLogger.SetOutput(fileHandler)
			lLogger.SetFormatter(getFormatter(config.FileJSONFormat))
		} else {
			lLogger.SetOutput(stdOutHandler)
		}
	}

	return &logrusLogger{
		logger: lLogger,
	}
}

// Debug - вывести отладку
func (l *logrusLogger) Debug(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info - вывести информацию
func (l *logrusLogger) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn - вывести предупреждение
func (l *logrusLogger) Warn(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// вывести ошибку
func (l *logrusLogger) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal - вывести критическую ошибку
func (l *logrusLogger) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Panic - вывести критическую ошибку
func (l *logrusLogger) Panic(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Printf - вывести текст
func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// WithFields - вывести дополнительные данные
func (l *logrusLogger) WithFields(fields Fields) Interface {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

// Debug - вывести отладку
func (l *logrusLogEntry) Debug(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Info - вывести информацию
func (l *logrusLogEntry) Info(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warn - вывести предупреждение
func (l *logrusLogEntry) Warn(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error - вывести ошибку
func (l *logrusLogEntry) Error(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatal - вывести критическую ошибку
func (l *logrusLogEntry) Fatal(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// Panic - вывести критическую ошибку
func (l *logrusLogEntry) Panic(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// Printf - вывести текст
func (l *logrusLogEntry) Printf(format string, args ...interface{}) {
	l.entry.Printf(format, args...)
}

// WithFields - Сохранить доп. поля
func (l *logrusLogEntry) WithFields(fields Fields) Interface {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

// convertToLogrusFields - конвертирует интерфейс в поля данных логгера
func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
