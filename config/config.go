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
		MorningHour: 8,
		MorningMin:  35,
		NoonHour:    15,
		NoonMin:     30,
		EveningHour: 20,
		EveningMin:  21,
	}
}
