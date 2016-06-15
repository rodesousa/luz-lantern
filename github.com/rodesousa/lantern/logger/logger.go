package logger

import (
	log "github.com/logrus"
	"io"
	"os"
)

// all Logger's type
const (
	Console LoggerType = iota
	Out
	File
	LogStash
)

const (
	NO_COLOR Color = 0
	RED = 31
	GREEN = 32
	YELLOW = 33
	BLUE = 34
	GRAY = 37
)

type Color uint8

// Log type
type LoggerType uint8

// Args type to log
type Fields map[string]interface{}

// Type Logger
type Logger struct {
	logger   *log.Logger
	log_type LoggerType
}

// AllLogger available
var AllLogger = []LoggerType{Console, Out, File, LogStash}

// logger Slice
// var loggers []*log.Logger
var loggers []Logger

func addLogger(log *log.Logger, log_type LoggerType) {
	// Add new logger
	loggers = append(loggers, Logger{log, log_type})
}

func GetOutLogger() io.Writer {
	return log.StandardLogger().Out
}

func Debug(msg string) {
	for _, logger := range loggers {
		logger.logger.Debug(msg)
	}
}

func DebugWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.logger.WithFields(log.Fields(args)).Debug(msg)
	}
}

func Info(msg string) {
	for _, logger := range loggers {
		logger.logger.Info(msg)
	}
}

func InfoColor(msg string, color Color) {
	for _, logger := range loggers {
		args := addColorToFields(logger.log_type, color, make(Fields))
		logger.logger.WithFields(log.Fields(args)).Info(msg)
	}
}

func InfoWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.logger.WithFields(log.Fields(args)).Info(msg)
	}
}

func InfoWithFieldsColor(msg string, args Fields, color Color) {
	for _, logger := range loggers {
		newargs := addColorToFields(logger.log_type, color, args)
		logger.logger.WithFields(log.Fields(newargs)).Info(msg)

	}
}

func WarnWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.logger.WithFields(log.Fields(args)).Warn(msg)
	}
}

func ErrorWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.logger.WithFields(log.Fields(args)).Error(msg)
	}
}

func FatalWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.logger.WithFields(log.Fields(args)).Fatal(msg)
	}
	os.Exit(1)
}

func PrintShardResult(msg string, testShard bool, cmdArg string, cmdOut string, cmdError error) {
	InfoWithFields(msg, Fields{"IsTestOk": testShard, "str_cmdArg": cmdArg, "str_error": cmdError, "str_out": cmdOut})
}

func addColorToFields(log_type LoggerType, color Color, args Fields) Fields {
	if (log_type == Console) {
		dst := make(Fields)
		for k, v := range args {
			dst[k] = v
		}
		dst["color"] = color
		return dst
	}
	return args
}

func Init(debug bool, toConsole bool, toFile bool, filenamePath string) {
	// Default logging level
	var logLevel = log.InfoLevel
	// if debug
	if debug {
		logLevel = log.DebugLevel
		log.SetLevel(log.DebugLevel)
	}
	// Init default logger
	log.SetLevel(logLevel)
	// If toConsole is set
	if toConsole {
		log.SetFormatter(ConsoleFormatter{})
		addLogger(log.StandardLogger(), Console)
	}
	// If toFile is set
	if toFile {
		f, err := os.OpenFile(filenamePath, os.O_APPEND | os.O_CREATE, 0755)
		if err != nil {
			WarnWithFields("Error creating or opening file", Fields{"filename": filenamePath})
		} else {
			tf := &log.TextFormatter{DisableColors: true}
			var logFile = &log.Logger{
				Out:       f,
				Formatter: tf,
				Level:     logLevel,
			}
			addLogger(logFile, File)
		}
	}
}
