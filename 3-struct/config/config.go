package config

import "os"

type Config struct {
	Key string
}

func GetConfig() *Config {
	key := os.Getenv("KEY")


	if key == "" {
		panic("You don't have KEY param")
	}

	return &Config{
		Key: key,
	}
}
