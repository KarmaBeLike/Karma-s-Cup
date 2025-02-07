package config

import (
	"time"
)

type Config struct {
	TimeZone    *time.Location
	MorningTime time.Time
	NoonTime    time.Time
	EveningTime time.Time
}

func New() *Config {
	tz, _ := time.LoadLocation("Asia/Atyrau")
	return &Config{
		TimeZone:    tz,
		MorningTime: time.Date(0, 0, 0, 18, 16, 0, 0, tz),
		NoonTime:    time.Date(0, 0, 0, 18, 19, 0, 0, tz),
		EveningTime: time.Date(0, 0, 0, 18, 22, 0, 0, tz),
	}
}
