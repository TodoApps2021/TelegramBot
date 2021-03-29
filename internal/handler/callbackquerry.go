package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	callBackQuerryDeleteList = "delete_list"
	callBackQuerryDeleteItem = "delete_item"
	callBackQuerryShowItems  = "show_items"
	callBackQuerryDoneItem   = "done_item"
	callBackQuerryBackItem   = "back_item"
)

func (h *Handler) HandleCallBackQuery(cbq *tgbotapi.CallbackQuery) error {
	switch cbq.Data {
	case callBackQuerryDeleteList:
		return h.DeleteList(cbq)
	case callBackQuerryShowItems:
		return h.ShowItems(cbq)
	case callBackQuerryDeleteItem:
		return h.DeleteItem(cbq)
	case callBackQuerryBackItem:
		return h.BackItem(cbq)
	case callBackQuerryDoneItem:
		return h.DoneItem(cbq)
	default:
		return nil
	}
}
