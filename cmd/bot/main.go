package main

import (
	"time"

	"github.com/todoapps2021/telegrambot/internal/handler"
	"github.com/todoapps2021/telegrambot/internal/storage"

	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  "1670019972:AAEH6MGGYuSWlrEduZ0DpFqqHbocorSEjX8",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}
	strg := &storage.TestStorage{}
	h := handler.New(b, strg, strg)
	b.Handle("/ls", h.ListTasks())
	b.Start()
}
