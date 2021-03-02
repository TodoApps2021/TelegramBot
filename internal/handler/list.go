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
	doneButton := selector.Data("Done✅", "done_button_list")
	removeButton := selector.Data("Remove🗑️", "remove_button_list")
	returnButton := selector.Data("Return↩", "return_button_list")

	h.bot.Handle(&doneButton, func(c *telebot.Callback) {
		returnButton.Data = c.Data
		removeButton.Data = c.Data
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
		doneButton.Data = c.Data
		removeButton.Data = c.Data
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
			doneButton.Data = v.ID
			removeButton.Data = v.ID
			returnButton.Data = v.ID
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
