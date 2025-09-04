package config

type Config struct {
	DBpath string
}

func LoadConfig() *Config {
	return &Config{
		DBpath: "data/workova.db",
	}
}
