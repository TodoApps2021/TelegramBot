package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/todoapps2021/telegrambot/internal/service"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func New(bot *tgbotapi.BotAPI, service *service.Service) *Handler {
	return &Handler{bot: bot, service: service}
}
