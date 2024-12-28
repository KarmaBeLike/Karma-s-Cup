package main

import (
	"log"
	"os"
	"time"

	"github.com/KarmaBeLike/Karma-s-Cup/config"
	"github.com/KarmaBeLike/Karma-s-Cup/scheduler"
	"github.com/joho/godotenv"
	telebot "gopkg.in/telebot.v3"
)

func main() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получаем токен из переменных окружения
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN не установлен")
	}
	cfg := config.New()

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	recipient := &telebot.User{
		ID: 873244236, // chat_id человека
	}

	scheduler := scheduler.New(bot, recipient, cfg)

	// Запускаем планировщик в отдельной горутине
	go scheduler.Start()

	log.Printf("Бот запущен. Расписание: утро - %02d:%02d, день - %02d:%02d, вечер - %02d:%02d\n",
		cfg.MorningHour, cfg.MorningMin,
		cfg.NoonHour, cfg.NoonMin,
		cfg.EveningHour, cfg.EveningMin)
	bot.Start()
}
