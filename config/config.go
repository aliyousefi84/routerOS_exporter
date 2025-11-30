package config

import "os"

type Config struct {
	RoterAddr  string
	RouterUser string
	RouterPass string
	SrvAddr    string
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.RoterAddr = os.Getenv("ROUTEROS_ADDRESS")
	cfg.RouterUser = os.Getenv("ROUTEROS_USER")
	cfg.RouterPass = os.Getenv("ROUTEROS_PASSWORD")
	cfg.SrvAddr = os.Getenv("SERVER_ADDRESS")
	return cfg
}
