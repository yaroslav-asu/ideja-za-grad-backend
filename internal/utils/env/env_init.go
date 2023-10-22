package env

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var (
	DbUser     string
	DbPassword string
	DbName     string
)
var (
	RunningMode      string
	TelegramBotToken string
	AdminChatId      int64
)

func initLoggerEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		zap.L().Fatal("Failed to load .env.db file")
	}
	RunningMode = os.Getenv("RUNNING_MODE")
	TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	AdminChatId, err = strconv.ParseInt(os.Getenv("ADMIN_CHAT_ID"), 10, 64)
	if err != nil {
		zap.L().Fatal("Failed to parse admin chat id")
	}
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
