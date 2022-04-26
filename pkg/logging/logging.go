package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// Build Log hook (methods Fire and Levels mush have!)
type changeOutputLog struct {
	LogLevels []logrus.Level
	Writer    []io.Writer
}

func (c *changeOutputLog) Fire(logEntry *logrus.Entry) error {
	// will be called every time.
	line, err := logEntry.String()
	if err != nil {
		return err
	}
	for _, wr := range c.Writer {
		if _, err := wr.Write([]byte(line)); err != nil {
			return err
		}
	}
	return nil
}

func (c *changeOutputLog) Levels() []logrus.Level {
	return c.LogLevels
}

func init() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	// logger.SetLevel(logrus.ErrorLevel) // use if there is no log hook.
	logger.SetOutput(io.Discard) // blocking the issuance of logs.
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)

			// func it's in, name file: line number
			return frame.Function, fmt.Sprintf("%s: %d", fileName, frame.Line)
		},
		DisableColors: true, // color on/off
		FullTimestamp: true, // time format
	}
	makeLogFolder()
	addLoggerHooks(logger)

	logEntry = logrus.NewEntry(logger) // HACK
}

func makeLogFolder() {
	// Create log folder
	if err := os.Mkdir("logs", 0600); !os.IsExist(err) {
		panic(err)
	}
}

func addLoggerHooks(logger *logrus.Logger) {
	files, err := os.OpenFile("logs/general.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger.AddHook(&changeOutputLog{
		LogLevels: logrus.AllLevels,
		Writer:    []io.Writer{files, os.Stdout},
	})

	// Alternative to Hooks (not recommended):
	// (to replace -> "logger.AddHook")
	// wrt := io.MultiWriter(os.Stdout, files)
	// logger.SetOutput(wrt)
	logger.SetLevel(logrus.TraceLevel) // watch all
}

///// HACK /////
var logEntry *logrus.Entry

type LogrusLogger struct {
	*logrus.Entry
}

func (l *LogrusLogger) LoggerWithField(key string, value interface{}) *LogrusLogger {
	return &LogrusLogger{l.WithField(key, value)}
}

func GetLogger() LogrusLogger {
	return LogrusLogger{logEntry}
}

///// HACK /////
