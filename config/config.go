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
		MorningMin:  35,
		NoonHour:    17,
		NoonMin:     20,
		EveningHour: 22,
		EveningMin:  21,
	}
}
