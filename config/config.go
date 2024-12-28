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
		MorningMin:  0,
		NoonHour:    16,
		NoonMin:     0,
		EveningHour: 23,
		EveningMin:  0,
	}
}
