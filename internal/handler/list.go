package handler

import (
	"bytes"
	"context"
	"log"
	"text/template"

	"gopkg.in/tucnak/telebot.v2"
)

const taskTemplateString = `Task #{{.ID}}: {{.Name}}
{{.Content}}`

const (
	doneButtonText   = "Done✅"
	removeButtonText = "Remove🗑️"
	returnButtonText = "Return↩"
	doneButtonID     = "id_list_done"
	removeButtonID   = "id_list_remove"
	returnButtonID   = "id_list_return"
)

var taskTemplate *template.Template

func init() {
	tmpl, err := template.New("taskTemplate").Parse(taskTemplateString)
	if err != nil {
		panic(err)
	}
	taskTemplate = tmpl
}

func (h *Handler) ListTasks() func(m *telebot.Message) {
	selector := &telebot.ReplyMarkup{}
	doneButton := selector.Data(doneButtonText, doneButtonID)
	removeButton := selector.Data(removeButtonText, removeButtonID)
	returnButton := selector.Data(returnButtonText, returnButtonID)

	h.bot.Handle(&doneButton, func(c *telebot.Callback) {
		selector := &telebot.ReplyMarkup{}
		removeButton := selector.Data(removeButtonText, removeButtonID, c.Data)
		returnButton := selector.Data(returnButtonText, returnButtonID, c.Data)
		selector.Inline(selector.Row(returnButton, removeButton))
		if _, err := h.bot.EditReplyMarkup(c.Message, selector); err != nil { // TODO: add error handling
			return
		}
		if err := h.bot.Respond(c, &telebot.CallbackResponse{
			Text:      "Good job, you've done your task!👍",
			ShowAlert: true,
		}); err != nil { // TODO: add error handling
			return
		}
	})
	h.bot.Handle(&removeButton, func(c *telebot.Callback) {
		log.Println("DELETE " + c.Data)
		if err := h.bot.Respond(c, &telebot.CallbackResponse{
			Text:      "The task's removed!🗑",
			ShowAlert: true,
		}); err != nil { // TODO: add error handling
			return
		}
		if err := h.bot.Delete(c.Message); err != nil {
			return
		}
	})
	h.bot.Handle(&returnButton, func(c *telebot.Callback) {
		selector := &telebot.ReplyMarkup{}
		doneButton := selector.Data(doneButtonText, doneButtonID, c.Data)
		removeButton := selector.Data(removeButtonText, removeButtonID, c.Data)
		selector.Inline(selector.Row(doneButton, removeButton))
		if _, err := h.bot.EditReplyMarkup(c.Message, selector); err != nil {
			return
		}
		if err := h.bot.Respond(c, &telebot.CallbackResponse{
			Text:      "No problem, it's returned now!👍",
			ShowAlert: true,
		}); err != nil { // TODO: add error handling
			return
		}
	})

	return func(m *telebot.Message) {
		tasks, err := h.reader.GetTasks(context.TODO(), 5, 0)
		if err != nil {
			log.Println(err)
			return
		}
		for _, v := range tasks {
			var buf bytes.Buffer
			if err = taskTemplate.Execute(&buf, v); err != nil {
				return
			}
			selector := &telebot.ReplyMarkup{}
			doneButton := selector.Data(doneButtonText, doneButtonID, v.ID)
			removeButton := selector.Data(removeButtonText, removeButtonID, v.ID)
			returnButton := selector.Data(returnButtonText, returnButtonID, v.ID)

			if v.Done {
				selector.Inline(selector.Row(returnButton, removeButton))
			} else {
				selector.Inline(selector.Row(doneButton, removeButton))
			}
			if _, err = h.bot.Send(m.Sender, buf.String(), selector); err != nil {
				return
			}
		}
	}
}
