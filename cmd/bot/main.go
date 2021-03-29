package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/todoapps2021/telegrambot/internal/handler"
	"github.com/todoapps2021/telegrambot/internal/kafka/producer"
	"github.com/todoapps2021/telegrambot/internal/repository"
	"github.com/todoapps2021/telegrambot/internal/service"
	"github.com/todoapps2021/telegrambot/internal/telegram"
)

func init() {
	err := godotenv.Load("../TelegramBot/.env")
	if err != nil {
		logrus.Fatal("error load env")
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		logrus.Fatal("error telegram token")
	}

	// bot.Debug = true // for logs

	postgres, _ := repository.NewPostgresDB(repository.Config{
		DB_URL: os.Getenv("DB_URL"),
	})
	kp, _ := producer.NewProducerKafka(producer.Config{
		Url: os.Getenv("KAFKA_URL"),
	})
	repository := repository.NewRepository(postgres)
	kafkaProducer := producer.NewKProducer(kp)
	service := service.NewService(repository, kafkaProducer)
	handler := handler.New(bot, service)

	tbot := telegram.NewBot(bot, handler)

	if err := tbot.Start(); err != nil {
		logrus.Fatal(err.Error())
	}
}
