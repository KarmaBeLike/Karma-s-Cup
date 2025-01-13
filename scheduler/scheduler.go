package scheduler

import (
	"log"
	"time"

	"github.com/KarmaBeLike/Karma-s-Cup/config"
	"github.com/KarmaBeLike/Karma-s-Cup/content"
	telebot "gopkg.in/telebot.v3"
)

type Scheduler struct {
	bot       *telebot.Bot
	chat      *telebot.Chat
	config    *config.Config
	location  *time.Location
	horoscope content.ContentProvider
	joke      content.ContentProvider
	quote     content.ContentProvider
}

func New(bot *telebot.Bot, config *config.Config) *Scheduler {
	loc, err := time.LoadLocation("Kazakhstan/Almaty")
	if err != nil {
		log.Println("failed to set time zone")
		loc = time.UTC
	}
	return &Scheduler{
		bot:       bot,
		config:    config,
		location:  loc,
		horoscope: content.NewHoroscopeProvider("aquarius"),
		joke:      content.NewJokeProvider(),
		quote:     content.NewQuoteProvider(),
	}
}

func (s *Scheduler) SetChat(chat *telebot.Chat) {
	s.chat = chat
	log.Printf("Установлен чат для отправки сообщений: %s (ID: %d)", chat.Title, chat.ID)
}

func (s *Scheduler) Start() {
	if s.chat == nil {
		log.Println("Чат не установлен! Ожидаем добавления бота в группу...")
		return
	}

	log.Printf("Планировщик запущен для группы: %s", s.chat.Title)
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		now := time.Now().In(s.location)
		hour := now.Hour()
		minute := now.Minute()

		switch {
		case hour == s.config.MorningHour && minute == s.config.MorningMin:
			s.sendMorningMessage()
		case hour == s.config.NoonHour && minute == s.config.NoonMin:
			s.sendNoonMessage()
		case hour == s.config.EveningHour && minute == s.config.EveningMin:
			s.sendEveningMessage()
		}
	}
}

func (s *Scheduler) sendMorningMessage() {
	if s.chat == nil {
		return
	}
	content, err := s.horoscope.GetContent()
	if err != nil {
		log.Printf("Ошибка получения гороскопа: %v", err)
		return
	}

	_, err = s.bot.Send(s.chat, content)
	if err != nil {
		log.Printf("Ошибка отправки утреннего сообщения: %v", err)
	}
}

func (s *Scheduler) sendNoonMessage() {
	if s.chat == nil {
		return
	}
	content, err := s.joke.GetContent()
	if err != nil {
		log.Printf("Ошибка получения анекдота: %v", err)
		return
	}

	_, err = s.bot.Send(s.chat, content)
	if err != nil {
		log.Printf("Ошибка отправки дневного сообщения: %v", err)
	}
}

func (s *Scheduler) sendEveningMessage() {
	if s.chat == nil {
		return
	}
	content, err := s.quote.GetContent()
	if err != nil {
		log.Printf("Ошибка получения цитаты: %v", err)
		return
	}

	_, err = s.bot.Send(s.chat, content)
	if err != nil {
		log.Printf("Ошибка отправки вечернего сообщения: %v", err)
	}
}
