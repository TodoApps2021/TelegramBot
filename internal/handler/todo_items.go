package handler

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) ShowItems(cbq *tgbotapi.CallbackQuery) error {
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

	items, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		text := "Todo lists is empty. Create on site"
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)
		// add button url
		if _, e := h.bot.Send(msg); e != nil {
			return e
		}

		return nil
	}

	text := fmt.Sprintf("Show items from list ID=%v", arr[1])
	msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)
	// add button url
	if _, e := h.bot.Send(msg); e != nil {
		return e
	}

	for _, v := range items {
		var status string
		if v.Done {
			status = "✅"
		} else {
			status = "❌"
		}
		text := fmt.Sprintf("ID: %v\nTitle: %v\nDescription: %v\nStatus: %v", v.Id, v.Title, v.Description, status)
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)

		if v.Done {
			msg.ReplyMarkup = INLINE_KEYBOARD_ITEM_2
		} else {
			msg.ReplyMarkup = INLINE_KEYBOARD_ITEM_1
		}

		if _, e := h.bot.Send(msg); e != nil {
			return e
		}
	}

	return nil
}

func (h *Handler) DeleteItem(cbq *tgbotapi.CallbackQuery) error {
	userId, _ := h.service.Authorization.GetUserIdByTgId(cbq.From.ID)
	if userId == 0 {
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	arr := strings.Split(cbq.Message.Text, "\n")
	arr1 := strings.Split(arr[0], ": ")

	itemId, err := strconv.Atoi(arr1[1])
	if err != nil {
		return err
	}

	if h.service.TodoItem.Delete(userId, itemId) != nil {
		return err
	}

	msg := tgbotapi.NewDeleteMessage(cbq.Message.Chat.ID, cbq.Message.MessageID)
	if _, e := h.bot.Send(msg); e != nil {
		return e
	}

	return nil
}

func (h *Handler) BackItem(cbq *tgbotapi.CallbackQuery) error {
	userId, _ := h.service.Authorization.GetUserIdByTgId(cbq.From.ID)
	if userId == 0 {
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	arr := strings.Split(cbq.Message.Text, "\n")
	arr1 := strings.Split(arr[0], ": ")

	itemId, err := strconv.Atoi(arr1[1])
	if err != nil {
		return err
	}

	if h.service.TodoItem.Update(userId, itemId, false) != nil {
		return err
	}

	text := strings.Replace(cbq.Message.Text, "✅", "❌", 1)
	msg := tgbotapi.NewEditMessageText(cbq.Message.Chat.ID, cbq.Message.MessageID, text)
	msg.ReplyMarkup = &INLINE_KEYBOARD_ITEM_1
	if _, e := h.bot.Send(msg); e != nil {
		return e
	}

	return nil
}

func (h *Handler) DoneItem(cbq *tgbotapi.CallbackQuery) error {
	userId, _ := h.service.Authorization.GetUserIdByTgId(cbq.From.ID)
	if userId == 0 {
		text := "You are anonymous. Go to site and register."
		msg := tgbotapi.NewMessage(cbq.Message.Chat.ID, text)

		_, err := h.bot.Send(msg)

		return err
	}

	arr := strings.Split(cbq.Message.Text, "\n")
	arr1 := strings.Split(arr[0], ": ")

	itemId, err := strconv.Atoi(arr1[1])
	if err != nil {
		return err
	}

	if h.service.TodoItem.Update(userId, itemId, true) != nil {
		return err
	}

	text := strings.Replace(cbq.Message.Text, "❌", "✅", 1)
	msg := tgbotapi.NewEditMessageText(cbq.Message.Chat.ID, cbq.Message.MessageID, text)
	msg.ReplyMarkup = &INLINE_KEYBOARD_ITEM_2
	if _, e := h.bot.Send(msg); e != nil {
		return e
	}

	return nil
}
