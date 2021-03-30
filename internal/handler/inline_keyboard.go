package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	INLINE_KEYBOARD_LISTS = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete 🔥", "delete_list"),
			tgbotapi.NewInlineKeyboardButtonData("Show items 📤", "show_items"),
		),
	)
	INLINE_KEYBOARD_ITEM_1 = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete 🔥", "delete_item"),
			tgbotapi.NewInlineKeyboardButtonData("Done ✅", "done_item"),
		),
	)
	INLINE_KEYBOARD_ITEM_2 = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete 🔥", "delete_item"),
			tgbotapi.NewInlineKeyboardButtonData("Back ⤴️", "back_item"),
		),
	)
)
