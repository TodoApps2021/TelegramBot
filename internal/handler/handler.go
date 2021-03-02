package handler

import (
	"github.com/todoapps2021/telegrambot/internal/storage"
	"gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	reader  storage.Reader
	creator storage.Creator
	bot     *telebot.Bot
}

func New(bot *telebot.Bot, reader storage.Reader, creator storage.Creator) *Handler {
	return &Handler{
		bot:     bot,
		reader:  reader,
		creator: creator,
	}
}
