package internal

import (
	"github.com/yaroslav-asu/urban-map/internal/logger"
	"github.com/yaroslav-asu/urban-map/internal/utils/db"
	"github.com/yaroslav-asu/urban-map/internal/utils/env"
)

func Init() {
	env.Init()
	logger.Init()
	db.Init()
}
