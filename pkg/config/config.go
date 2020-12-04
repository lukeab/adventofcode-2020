package config

//Config settings for the application
type Config struct {
	Debug       bool
	Inputfile   string
	Targetvalue int //defaults to 0
}

//New generate and return config.Config struct
func New() (*Config, error) {
	cfg := new(Config)
	cfg.Debug = false
	//cfg.Inputfile = ""
	// cfg.Targetvalue = 0
	return cfg, nil
}
