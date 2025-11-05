package config

import "os"

type Config struct {
	RoterAddr string
	RouterUser string
	RouterPass string
	SrvAddr string
}

func  InitEnv () *Config {
	return &Config{
		RoterAddr: os.Getenv("ROUTEROS_ADDRESS"),
		RouterUser: os.Getenv("ROUTEROS_USER"),
		RouterPass: os.Getenv("ROUTEROS_PASSWORD"),
		SrvAddr: os.Getenv("SERVER_ADDRESS"),
	}
}