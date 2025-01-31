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
		MorningHour: 21,
		MorningMin:  35,
		NoonHour:    17,
		NoonMin:     50,
		EveningHour: 22,
		EveningMin:  14,
	}
}
