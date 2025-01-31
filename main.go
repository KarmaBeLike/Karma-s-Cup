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
		Token:   token,
		Poller:  &telebot.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New()
	scheduler := scheduler.New(bot, cfg)

	bot.Handle(telebot.OnAddedToGroup, func(c telebot.Context) error {
		log.Printf("Обработчик группы ВЫЗВАН")
		log.Printf("Полная информация о чате: %+v", c.Chat())

		if c.Chat() == nil {
			log.Println("ОШИБКА: Chat() вернул nil")
			return nil
		}

		log.Printf("ID чата: %v", c.Chat().ID)
		log.Printf("Название чата: %s", c.Chat().Title)
		log.Printf("Тип чата: %s", c.Chat().Type)

		scheduler.SetChat(c.Chat())
		go scheduler.Start()

		return c.Send("Good Morning, Vietnam!")
	})
}
