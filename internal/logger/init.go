package logger

import (
	"github.com/yaroslav-asu/urban-map/internal/utils/env"
	"go.uber.org/zap"
	"log"
)

func Init() {
	var logger *zap.Logger
	var err error
	switch env.RunningMode {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal("Failed to initialize logger")
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
