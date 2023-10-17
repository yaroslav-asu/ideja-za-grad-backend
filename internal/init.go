package internal

import (
	"urban-map/internal/logger"
	"urban-map/internal/utils/db"
	"urban-map/internal/utils/env"
)

func Init() {
	env.Init()
	logger.Init()
	db.Init()
}
