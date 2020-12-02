package config

type Config struct {
	Debug       bool
	Inputfile   string
	Targetvalue int
}

func New() (*Config, error) {
	cfg := new(Config)
	cfg.Debug = false
	cfg.Inputfile = "./input"
	cfg.Targetvalue = 2020
	return cfg, nil
}
