package logger

import (
	log "github.com/logrus"
	"io"
	"os"
)

// Args type to log
type Fields map[string]interface{}

// Variables
var logToConsole bool
var logToFile bool
var logFile *log.Logger

func Init(debug bool, toConsole bool, toFile bool, filenamePath string) {
	logToConsole = toConsole
	// Default logging level
	var logLevel = log.InfoLevel
	log.SetLevel(logLevel)
	// if debug
	if debug {
		logLevel = log.DebugLevel
		log.SetLevel(log.DebugLevel)
	}
	// If toFile is set
	if (toFile) {
		f, err := os.OpenFile(filenamePath, os.O_APPEND | os.O_CREATE, 0755)
		if err != nil {
			WarnWithFields("Error creating or opening file", Fields{"filename":filenamePath})
		} else {
			tf := &log.TextFormatter{DisableColors: true}
			logFile = &log.Logger{
				Out: f,
				Formatter: tf,
				Level: logLevel,
			}
			logToFile= true
		}
	}
}

func GetOutLogger() io.Writer {
	return log.StandardLogger().Out
}

func Debug(msg string) {
	if logToConsole {
		log.Debug(msg)
	}
	if logToFile {
		logFile.Debug(msg)
	}
}

func DebugWithFields(msg string, args Fields) {
	if logToConsole {
		log.WithFields(log.Fields(args)).Debug(msg)
	}
	if logToFile {
		logFile.WithFields(log.Fields(args)).Debug(msg)
	}
}

func Info(msg string) {
	if logToConsole {
		log.Info(msg)
	}
	if logToFile {
		logFile.Info(msg)
	}
}

func InfoWithFields(msg string, args Fields) {
	if logToConsole {
		log.WithFields(log.Fields(args)).Info(msg)
	}
	if logToFile {
		logFile.WithFields(log.Fields(args)).Info(msg)
	}
}

func WarnWithFields(msg string, args Fields) {
	if logToConsole {
		log.WithFields(log.Fields(args)).Warn(msg)
	}
	if logToFile {
		logFile.WithFields(log.Fields(args)).Warn(msg)
	}
}

func ErrorWithFields(msg string, args Fields) {
	if logToConsole {
		log.WithFields(log.Fields(args)).Error(msg)
	}
	if logToFile {
		logFile.WithFields(log.Fields(args)).Error(msg)
	}
}

func FatalWithFields(msg string, args Fields) {
	if logToConsole {
		log.WithFields(log.Fields(args)).Fatal(msg)
	}
	if logToFile {
		logFile.WithFields(log.Fields(args)).Fatal(msg)
	}
	os.Exit(1)
}

func PrintShardResult(msg string, testShard bool, cmdArg string, cmdOut string, cmdError error) {
	InfoWithFields(msg, Fields{"IsTestOk" : testShard, "str_cmdArg": cmdArg, "str_error" : cmdError, "str_out": cmdOut})
}