package logger

import (
	"log"
	"os"
)

type Logger struct {
	logInfo     *log.Logger
	logWarn     *log.Logger
	logErr      *log.Logger
	logFilePath string
}

func New(logFilePath string) (*Logger, error) {
	openFlags := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	fileInfo, err := os.OpenFile(logFilePath, openFlags, 0666)
	if err != nil {
		return nil, err
	}

	fileWarn, err := os.OpenFile(logFilePath, openFlags, 0666)
	if err != nil {
		return nil, err
	}

	fileErr, err := os.OpenFile(logFilePath, openFlags, 0666)
	if err != nil {
		return nil, err
	}

	logFlags := log.LstdFlags

	logInfo := log.New(fileInfo, "INFO:\t", logFlags)
	logWarn := log.New(fileWarn, "WARN:\t", logFlags)
	logErr := log.New(fileErr, "ERR:\t", logFlags)

	return &Logger{
		logInfo:     logInfo,
		logWarn:     logWarn,
		logErr:      logErr,
		logFilePath: logFilePath,
	}, nil
}

func (l *Logger) Info(text ...any) {
	l.logInfo.Println(text...)
}

func (l *Logger) Warn(text ...any) {
	l.logWarn.Println(text...)
}

func (l *Logger) Err(text ...any) {
	l.logErr.Println(text...)
}
