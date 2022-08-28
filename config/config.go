package config

import "os"

type Config struct {
	Port   string
	ConnDb string
}

func NewConfig() *Config {
	config := Config{}
	config.Port = os.Getenv("PORT")
	config.ConnDb = os.Getenv("CONN_DB")
	return &config
}
