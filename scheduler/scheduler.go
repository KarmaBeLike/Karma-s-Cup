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
	recipient *telebot.User
	config    *config.Config
	horoscope content.ContentProvider
	joke      content.ContentProvider
	quote     content.ContentProvider
}

func New(bot *telebot.Bot, recipient *telebot.User, config *config.Config) *Scheduler {
	return &Scheduler{
		bot:       bot,
		recipient: recipient,
		config:    config,
		horoscope: content.NewHoroscopeProvider("Aquarius"),
		joke:      content.NewJokeProvider(),
		quote:     content.NewQuoteProvider(),
	}
}

func (s *Scheduler) Start() {
	log.Println("Планировщик запущен")
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()

		switch {
		case hour == s.config.MorningHour && minute == s.config.MorningMin:
			log.Printf("Отправка утреннего сообщения в %02d:%02d", hour, minute)
			s.sendMorningMessage()
		case hour == s.config.NoonHour && minute == s.config.NoonMin:
			log.Printf("Отправка дневного сообщения в %02d:%02d", hour, minute)
			s.sendNoonMessage()
		case hour == s.config.EveningHour && minute == s.config.EveningMin:
			log.Printf("Отправка вечернего сообщения в %02d:%02d", hour, minute)
			s.sendEveningMessage()
		}
	}
}

func (s *Scheduler) sendMorningMessage() {
	content, err := s.horoscope.GetContent()
	if err != nil {
		log.Printf("Ошибка получения гороскопа: %v", err)
		return
	}

	_, err = s.bot.Send(s.recipient, "Гороскоп на сегодня:\n"+content)
	if err != nil {
		log.Printf("Ошибка отправки утреннего сообщения: %v", err)
	}
}

func (s *Scheduler) sendNoonMessage() {
	content, err := s.joke.GetContent()
	if err != nil {
		log.Printf("Ошибка получения анекдота: %v", err)
		return
	}

	_, err = s.bot.Send(s.recipient, "Анекдот дня:\n"+content)
	if err != nil {
		log.Printf("Ошибка отправки дневного сообщения: %v", err)
	}
}

func (s *Scheduler) sendEveningMessage() {
	content, err := s.quote.GetContent()
	if err != nil {
		log.Printf("Ошибка получения цитаты: %v", err)
		return
	}

	_, err = s.bot.Send(s.recipient, "Цитата дня:\n"+content)
	if err != nil {
		log.Printf("Ошибка отправки вечернего сообщения: %v", err)
	}
}
