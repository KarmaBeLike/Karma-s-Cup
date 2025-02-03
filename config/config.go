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
		MorningHour: 9,
		MorningMin:  59,
		NoonHour:    18,
		NoonMin:     15,
		EveningHour: 23,
		EveningMin:  14,
	}
}
