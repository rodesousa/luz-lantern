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

// Log type
type LoggerType uint8

// Args type to log
type Fields map[string]interface{}

// AllLogger available
var AllLogger = []LoggerType{Console, Out, File, LogStash}

// logger Slice
var loggers []*log.Logger

func addLogger(log *log.Logger) {
	// Add new logger
	loggers = append(loggers, log)
}

func GetOutLogger() io.Writer {
	return log.StandardLogger().Out
}

func Debug(msg string) {
	for _, logger := range loggers {
		logger.Debug(msg)
	}
}

func DebugWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.WithFields(log.Fields(args)).Debug(msg)
	}
}

func Info(msg string) {
	for _, logger := range loggers {
		logger.Info(msg)
	}
}

func InfoWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.WithFields(log.Fields(args)).Info(msg)
	}
}

func WarnWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.WithFields(log.Fields(args)).Warn(msg)
	}
}

func ErrorWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.WithFields(log.Fields(args)).Error(msg)
	}
}

func FatalWithFields(msg string, args Fields) {
	for _, logger := range loggers {
		logger.WithFields(log.Fields(args)).Fatal(msg)
	}
	os.Exit(1)
}

func PrintShardResult(msg string, testShard bool, cmdArg string, cmdOut string, cmdError error) {
	InfoWithFields(msg, Fields{"IsTestOk": testShard, "str_cmdArg": cmdArg, "str_error": cmdError, "str_out": cmdOut})
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
		addLogger(log.StandardLogger())
	}
	// If toFile is set
	if toFile {
		f, err := os.OpenFile(filenamePath, os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			WarnWithFields("Error creating or opening file", Fields{"filename": filenamePath})
		} else {
			tf := &log.TextFormatter{DisableColors: true}
			var logFile = &log.Logger{
				Out:       f,
				Formatter: tf,
				Level:     logLevel,
			}
			addLogger(logFile)
		}
	}
}
