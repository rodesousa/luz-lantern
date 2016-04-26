package logger

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
)


// Args type to log
type Fields map[string]interface{}

func Init(debug bool) {
	log.SetLevel(log.InfoLevel)
	if debug {
		log.SetLevel(log.DebugLevel)
	}

}

func GetOutLogger() io.Writer {
	return log.StandardLogger().Out
}

func Debug(msg string) {
	log.Debug(msg)
}

func DebugWithFields(msg string, args Fields) {
	log.WithFields(log.Fields(args)).Debug(msg)
}

func Info(msg string) {
	log.Info(msg)
}

func InfoWithFields(msg string, args Fields) {
	log.WithFields(log.Fields(args)).Info(msg)
}

func Warn(msg string, args Fields) {
	log.WithFields(log.Fields(args)).Warn(msg)
}

func Error(msg string, args Fields) {
	log.WithFields(log.Fields(args)).Error(msg)
}

func Fatal(msg string, args Fields) {
	log.WithFields(log.Fields(args)).Error(msg)
	os.Exit(1)
}

func ByteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n + 1])
}