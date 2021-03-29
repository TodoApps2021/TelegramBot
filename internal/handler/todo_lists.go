package handler

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/todoapps2021/telegrambot/internal/models"
)

func (h *Handler) GetAllList(message *tgbotapi.Message, userId int) error {
	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		return err
	}

	if len(lists) == 0 {
		text := "Todo lists is empty. Create on site"
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		// add button url
		if _, e := h.bot.Send(msg); e != nil {
			return e
		}

		return nil
	}

	for _, v := range lists {
		text := fmt.Sprintf("ID: %v\nTitle: %v\nDescription: %v", v.Id, v.Title, v.Description)
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		msg.ReplyMarkup = INLINE_KEYBOARD_LISTS
		if _, e := h.bot.Send(msg); e != nil {
			return e
		}
	}

	return nil
}

func (h *Handler) CreateList(message *tgbotapi.Message, userId int) error {
	arr := strings.SplitAfterN(message.Text, " ", 2)
	arr = strings.Split(arr[1], " : ")
	if len(arr) != 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Incorrect arg.")
		msg.ReplyToMessageID = message.MessageID
		if _, e := h.bot.Send(msg); e != nil {
			return e
		}
		return nil
	}

	var list models.TodoList

	list.Title = arr[0]
	list.Description = arr[1]

	return h.service.TodoList.Create(userId, list)
}

func (h *Handler) DeleteList(cbq *tgbotapi.CallbackQuery) error {
	userId, _ := h.service.Authorization.GetUserIdByTgId(cbq.From.ID)
	if userId == 0 {
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	arr := strings.Split(cbq.Message.Text, "\n")
	arr = strings.Split(arr[0], ": ")

	listId, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}

	if h.service.TodoList.Delete(userId, listId) != nil {
		return err
	}

	msg := tgbotapi.NewDeleteMessage(cbq.Message.Chat.ID, cbq.Message.MessageID)
	if _, e := h.bot.Send(msg); e != nil {
		return e
	}

	return nil
}
