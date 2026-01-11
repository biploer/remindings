package config

import (
	"os"

	"github.com/biploer/remindings/internal/adapter/postgres"
)

type config struct {
	BotToken string
	Postgres postgres.Config
}

func MustConfig() config {
	botToken := os.Getenv("BOT_TOKEN")

	if botToken == "" {
		panic("BOT_TOKEN env variable is must be")
	}

	return config{
		BotToken: botToken,
		Postgres: postgres.Config{},
	}
}
