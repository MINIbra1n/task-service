package config

import "time"

type AppConfig struct {
	LogLevel string
	Rest     Rest
}

type Rest struct {
	ListenAddress string        `envconfig:"PORT" required:"true"`
	WriteTimeout  time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	ServerName    string        `envconfig:"SERVER_NAME" required:"true"`
}
