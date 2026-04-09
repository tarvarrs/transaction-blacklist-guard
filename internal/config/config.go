package config

import "os"

type Config struct {
	HTTPAddr     string
	PGConnString string
}

func Load() Config {
	return Config{
		HTTPAddr:     ":8080",
		PGConnString: os.Getenv("POSTGRES_CONNECTION_STRING"),
	}
}
