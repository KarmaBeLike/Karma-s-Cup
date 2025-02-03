package config

import "time"

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
		MorningTime: time.Date(0, 0, 0, 9, 59, 0, 0, tz),
		NoonTime:    time.Date(0, 0, 0, 18, 36, 0, 0, tz),
		EveningTime: time.Date(0, 0, 0, 23, 14, 0, 0, tz),
	}
}
