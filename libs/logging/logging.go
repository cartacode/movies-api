package logging

import (
	"fmt"

	"go.uber.org/zap"
)

// FATAL --
const (
	FATAL   = 50
	ERROR   = 40
	WARNING = 30
	INFO    = 20
	DEBUG   = 10
	NOTSET  = 0
)

// atom -- Change the log level
var atom = zap.NewAtomicLevel()

// GetLog --
func GetLog() *zap.SugaredLogger {

	cfg := zap.NewDevelopmentConfig()

	cfg.OutputPaths = []string{
		"stdout",
	}
	cfg.Level = atom
	logger, err := cfg.Build()
	if err != nil {
		error.Error(err)
	}
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	return log
}

// GetProdLog --
func GetProdLog() *zap.SugaredLogger {

	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{
		"stdout",
	}
	cfg.Level = atom
	logger, err := cfg.Build()
	if err != nil {
		error.Error(err)
	}
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	return log
}

// SetLevel -- Set the Level for Logging -- Like Python
func SetLevel(level int) {
	switch level {
	case DEBUG:
		atom.SetLevel(zap.DebugLevel)
	case INFO:
		atom.SetLevel(zap.InfoLevel)
	case ERROR:
		atom.SetLevel(zap.ErrorLevel)
	case WARNING:
		atom.SetLevel(zap.WarnLevel)
	case FATAL:
		atom.SetLevel(zap.FatalLevel)
	default:
		fmt.Println("Log Level Not Found", level)
	}

}
