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
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN не установлен")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New()
	scheduler := scheduler.New(bot, cfg)

	bot.Handle(telebot.OnAddedToGroup, func(c telebot.Context) error {
		log.Printf("Бот добавлен в группу: %s", c.Chat().Title)
		scheduler.SetChat(c.Chat())
		go scheduler.Start()
		return c.Send("Good Morning, Vietnam!")
	})

	log.Println("Бот запущен. Добавьте его в группу.")
	bot.Start()
}
