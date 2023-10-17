package main

import (
	"go.uber.org/zap"
	"urban-map/internal"
	"urban-map/internal/utils/db"
)

func main() {
	internal.Init()
	d := db.Connect()
	defer db.Close(d)
	zap.L().Info("DB connection wasn't close: ")

}
