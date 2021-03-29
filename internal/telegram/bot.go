package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/todoapps2021/telegrambot/internal/handler"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	handler handler.Handler
}

func NewBot(bot *tgbotapi.BotAPI, handler *handler.Handler) *Bot {
	return &Bot{bot: bot, handler: *handler}
}

func (b *Bot) Start() error {
	logrus.Infof("Authorized on account %s (http://t.me/%s)", b.bot.Self.UserName, b.bot.Self.UserName)
	updates, err := b.getUpdatesChannel()
	if err != nil {
		return err
	}

	if err := b.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (b *Bot) getUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.CallbackQuery != nil {
			if err := b.handler.HandleCallBackQuery(update.CallbackQuery); err != nil {
				return err
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handler.HandleCommand(update.Message); err != nil {
				return err
			}
			continue
		}

		if err := b.handler.HandleMessage(update.Message); err != nil {
			return err
		}
	}

	return nil
}
