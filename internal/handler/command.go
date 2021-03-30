package handler

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const (
	commandStart  = "start"
	commandHelp   = "help"
	commandList   = "list"
	commandLogin  = "login"
	commandCreate = "create"
)

func (h *Handler) HandleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return h.Start(message)
	case commandCreate:
		return h.Create(message)
	case commandList:
		return h.List(message)
	case commandHelp:
		return h.Help(message)
	case commandLogin:
		return h.Login(message)
	default:
		return h.UnknownCommand(message)
	}
}

func (h *Handler) Start(message *tgbotapi.Message) error {
	text := fmt.Sprintf("Hello @%s.\nYou has telegram id = %v\nWrite /login <username>, <password>. \n\nAfter login I remeber your telegram id.\nHave a nice day!", message.From.UserName, message.From.ID)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	_, err := h.bot.Send(msg)

	defer func() {
		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := h.bot.Send(msg)
		if err != nil {
			return
		}
	}()

	return err
}

func (h *Handler) Create(message *tgbotapi.Message) error {
	userId, _ := h.service.Authorization.GetUserIdByTgId(message.From.ID)
	if userId == 0 {
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	defer func() {
		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := h.bot.Send(msg)
		if err != nil {
			return
		}
	}()

	return h.CreateList(message, userId)
}

func (h *Handler) Help(message *tgbotapi.Message) error { //TODO
	return nil
}

func (h *Handler) List(message *tgbotapi.Message) error {
	userId, err := h.service.Authorization.GetUserIdByTgId(message.From.ID)
	if userId == 0 {
		logrus.Info(userId, message.From.ID, err)
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	defer func() {
		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := h.bot.Send(msg)
		if err != nil {
			return
		}
	}()

	return h.GetAllList(message, userId)
}

func (h *Handler) Login(message *tgbotapi.Message) error {
	text := message.Text
	words := strings.Split(text, " ")
	if len(words) != 3 {
		logrus.Error("error")
	}

	if err := h.service.Authorization.Login(words[1], words[2], message.From.ID); err != nil {
		text := "Invalid username or password. Maybe you aren't registered?"
		msg := tgbotapi.NewMessage(message.Chat.ID, text)

		_, err := h.bot.Send(msg)
		return err
	}

	defer func() {
		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := h.bot.Send(msg)
		if err != nil {
			return
		}
	}()

	text = "Success\\! 🎉\n_Write command to get_ '/list'"
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "MarkdownV2"
	_, err := h.bot.Send(msg)

	return err
}

func (h *Handler) UnknownCommand(message *tgbotapi.Message) error {
	text := fmt.Sprintf("You write '%s'\nI didn't know this command.", message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	_, err := h.bot.Send(msg)

	defer func() {
		msg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		_, err := h.bot.Send(msg)
		if err != nil {
			return
		}
	}()

	return err
}
