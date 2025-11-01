package config

import "os"

type Config struct {
	Addr string
	User string
	Pass string
}

func  InitEnv () *Config {
	return &Config{
		Addr: os.Getenv("ROUTEROS_ADDRESS"),
		User: os.Getenv("ROUTEROS_USER"),
		Pass: os.Getenv("ROUTEROS_PASSWORD"),
	}
}