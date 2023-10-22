package telegram_bot

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"sync"
	"urban-map/internal/utils/db"
	"urban-map/internal/utils/env"
)

var tgBot *tgbotapi.BotAPI

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(env.TelegramBotToken)
	if err != nil {
		zap.L().Error("failed to init telegram Bot: " + err.Error())
	}
	switch env.RunningMode {
	case "dev":
		bot.Debug = true
	case "prod":
		bot.Debug = false
	default:
		bot.Debug = false
		zap.L().Warn("unknown running mode: " + env.RunningMode)
	}
	return bot
}
func removeMarkup(messageId int) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{})
	_, err := tgBot.Request(tgbotapi.NewEditMessageReplyMarkup(env.AdminChatId, messageId, keyboard))
	if err != nil {
		zap.L().Warn("failed to remove markup: " + err.Error())
	}
}
func Run(wg *sync.WaitGroup) {
	tgBot = Init()
	defer wg.Done()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := tgBot.Request(callback); err != nil {
				panic(err)
			}
			switch update.CallbackQuery.Data {
			case "approve":
				var m MarkerNotification
				err := json.Unmarshal([]byte(update.CallbackQuery.Message.Text), &m)
				if err != nil {
					zap.L().Error("failed to unmarshal message: " + err.Error())
				}
				markerFromNotify := m.Marker()
				err = markerFromNotify.Approve(db.GetDB())
				if err != nil {
					zap.L().Error(err.Error())
				}
				removeMarkup(update.CallbackQuery.Message.MessageID)
			case "decline":
				var m MarkerNotification
				err := json.Unmarshal([]byte(update.CallbackQuery.Message.Text), &m)
				if err != nil {
					zap.L().Warn("failed to unmarshal message: " + err.Error())
				}
				markerFromNotify := m.Marker()
				_ = markerFromNotify.Delete(db.GetDB())
				removeMarkup(update.CallbackQuery.Message.MessageID)
			default:
				zap.L().Warn("unknown callback data: " + update.CallbackQuery.Data)
			}
		}
	}
}
