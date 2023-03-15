package logger

// Fields - интерфейс набора дополнительных полей
type Fields map[string]interface{}

const (
	// Debug Вывод отладочной информации
	Debug = "debug"
	// Info Вывод логов по умолчанию
	Info = "info"
	// Warn Вывод сообщений о возможных проблемах
	Warn = "warn"
	// Error Вывод ошибок
	Error = "error"
	// Fatal Вывод критических ошибок. Приложение завершает работу после регистрации сообщения.
	Fatal = "fatal"
)

// Interface - интерфейс логгера
type Interface interface {
	// Debug Вывести отладку
	Debug(format string, args ...interface{})
	// Info Вывести информацию
	Info(format string, args ...interface{})
	// Warn Вывести предупреждение
	Warn(format string, args ...interface{})
	// Error Вывести ошибку
	Error(format string, args ...interface{})
	// Fatal Вывести критическую ошибку
	Fatal(format string, args ...interface{})
	// Panic Вывести критическую ошибку
	Panic(format string, args ...interface{})
	// WithFields Вывести дополнительные данные
	WithFields(keyValues Fields) Interface
	// Printf Вывести текст
	Printf(format string, args ...interface{})
}

// Configuration - объект конфигурации логгера
type Configuration struct {
	EnableConsole      bool   // Разрешить вывод в консоль
	ConsoleJSONFormat  bool   // Формат вывода в консоль
	ConsoleLevel       string // Уровень ошибок для вывода в консоль
	EnableFile         bool   // Разрешить сохранять в файл
	FileJSONFormat     bool   // Формат сохранения в файл
	FileLevel          string // Уровень ошибок для сохранения в файл
	FileLocation       string // Путь до папки с логами
	FileLocationPrefix string // Префикс названия файла логов
}
