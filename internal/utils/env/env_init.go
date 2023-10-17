package env

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

var (
	DbUser     string
	DbPassword string
	DbName     string
)
var (
	RunningMode string
)

func initLoggerEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		zap.L().Fatal("Failed to load .env.db file")
	}
	RunningMode = os.Getenv("RUNNING_MODE")
}

func initDbEnv() {
	err := godotenv.Load(".env.db")
	if err != nil {
		zap.L().Fatal("Failed to load .env.db file")
	}
	DbUser = os.Getenv("POSTGRES_USER")
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
}

func Init() {
	initLoggerEnv()
	initDbEnv()
}
