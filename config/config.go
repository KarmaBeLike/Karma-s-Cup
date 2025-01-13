package config

type Config struct {
	MorningHour int
	MorningMin  int
	NoonHour    int
	NoonMin     int
	EveningHour int
	EveningMin  int
}

func New() *Config {
	return &Config{
		MorningHour: 10,
		MorningMin:  37,
		NoonHour:    19,
		NoonMin:     05,
		EveningHour: 23,
		EveningMin:  59,
	}
}
