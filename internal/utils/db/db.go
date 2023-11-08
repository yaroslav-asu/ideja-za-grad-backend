package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"time"
	"urban-map/internal/utils/env"
	"urban-map/models/gorm/marker"
)

var reconnectTime = 5 * time.Second
var DB *gorm.DB

func Init() {
	db := Connect()
	defer Close(db)
	err := db.AutoMigrate(
		//&marker.Image{},
		&marker.Marker{},
		//&marker.Type{},
	)
	typesTitles := []string{"accessibility", "architecture", "bike_infrastructure", "pets_infrastructure", "playgrounds", "transport", "walkability", "natural"}
	for i, t := range typesTitles {
		db.Model(&marker.Type{}).FirstOrCreate(&marker.Type{ID: uint(i), Title: t})
	}
	if err != nil {
		zap.L().Error("failed to auto migrate database")
		zap.L().Info("Continuing without auto migration")
	}
}

func Connect() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", env.DbUser, url.QueryEscape(env.DbPassword), env.DbHost, env.DbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		zap.L().Error("Failed to connect db")
		zap.L().Info("Trying to reconnect db")
		time.Sleep(reconnectTime)
		reconnectTime *= 2
		return Connect()
	}
	DB = db
	return db
}

func Close(db *gorm.DB) {
	postgresDB, err := db.DB()
	if err != nil {
		zap.L().Error("Failed to get db instance: " + err.Error())
		zap.L().Info("DB connection wasn't close")
		return
	}
	err = postgresDB.Close()
	if err != nil {
		zap.L().Info("DB connection wasn't close: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return DB
}
