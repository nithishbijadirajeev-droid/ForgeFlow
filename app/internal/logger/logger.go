package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	log  *zap.Logger
	once sync.Once
)

// GetLogger returns a singleton Zap logger.
func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		log, err = zap.NewProduction()

		if err != nil {
			panic(err)
		}
	})

	return log
}

// Sync flushes any buffered log entries.
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
