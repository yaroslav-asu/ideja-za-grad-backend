package telegram_bot

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"urban-map/internal/utils/env"
	"urban-map/models/gorm/marker"
)

type MarkerNotification struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

func (m *MarkerNotification) Marker() marker.Marker {
	return marker.Marker{
		Type:        marker.Type{Title: m.Type},
		Description: m.Description,
		Coords: marker.Coords{
			Lat: m.Lat,
			Lng: m.Lng,
		},
	}
}

func SendNotification(m *marker.Marker) {
	lat, lng := m.Coords.Lat, m.Coords.Lng
	_, err := tgBot.Send(tgbotapi.NewLocation(env.AdminChatId, lat, lng))
	if err != nil {
		zap.L().Warn("failed to send message to admin: " + err.Error())
	}
	r := MarkerNotification{
		Type:        m.Type.Title,
		Description: m.Description,
		Lat:         lat,
		Lng:         lng,
	}
	j, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		zap.L().Warn("failed to marshal marker: " + err.Error())
	}
	message := tgbotapi.NewMessage(env.AdminChatId, string(j))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Approve", "approve"),
			tgbotapi.NewInlineKeyboardButtonData("Decline", "decline"),
		),
	)
	message.ReplyMarkup = &keyboard
	_, err = tgBot.Send(message)
	if err != nil {
		zap.L().Warn("failed to send message to admin: " + err.Error())
	}

}
