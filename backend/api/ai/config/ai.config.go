package config

import (
	"os"
)

type Config struct {
	APIKey string
	Model string
}

func GetConfig() *Config {
	key := os.Getenv("OPENAI_KEY")
	model := os.Getenv("OPENAI_MODEL")
	return &Config{
		APIKey : key,
		Model : model,
	}
}