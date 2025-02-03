package scheduler

import (
	"log"
	"time"

	"github.com/KarmaBeLike/Karma-s-Cup/config"
	"github.com/KarmaBeLike/Karma-s-Cup/content"
	telebot "gopkg.in/telebot.v3"
)

type Scheduler struct {
	bot          *telebot.Bot
	chat         *telebot.Chat
	config       *config.Config
	providers    map[string]content.ContentProvider
	sentMessages map[string]bool
}

func New(bot *telebot.Bot, cfg *config.Config) *Scheduler {
	return &Scheduler{
		bot:    bot,
		config: cfg,
		providers: map[string]content.ContentProvider{
			"morning": content.NewHoroscopeProvider("aquarius"),
			"noon":    content.NewJokeProvider(),
			"evening": content.NewQuoteProvider(),
		},
		sentMessages: make(map[string]bool),
	}
}

func (s *Scheduler) SetChat(chat *telebot.Chat) {
	s.chat = chat
	log.Printf("Чат установлен: %s (ID: %d)", chat.Title, chat.ID)
}

func (s *Scheduler) Start() {
	if s.chat == nil {
		log.Println("Чат не установлен. Ожидание добавления в группу...")
		return
	}

	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		s.checkAndSendMessages()
	}
}

func (s *Scheduler) checkAndSendMessages() {
	now := time.Now().In(s.config.TimeZone)
	currentTime := time.Date(0, 0, 0, now.Hour(), now.Minute(), 0, 0, time.UTC)

	s.checkMessage("morning", currentTime, s.config.MorningTime)
	s.checkMessage("noon", currentTime, s.config.NoonTime)
	s.checkMessage("evening", currentTime, s.config.EveningTime)
}

func (s *Scheduler) checkMessage(key string, currentTime, targetTime time.Time) {
	if !s.sentMessages[key] &&
		currentTime.Hour() == targetTime.Hour() &&
		currentTime.Minute() >= targetTime.Minute() &&
		currentTime.Minute() < targetTime.Minute()+2 {

		content, err := s.providers[key].GetContent()
		if err != nil {
			log.Printf("Ошибка получения %s контента: %v", key, err)
			return
		}

		_, err = s.bot.Send(s.chat, content)
		if err != nil {
			log.Printf("Ошибка отправки %s сообщения: %v", key, err)
			return
		}

		s.sentMessages[key] = true
	} else if currentTime.Hour() != targetTime.Hour() {
		s.sentMessages[key] = false
	}
}
