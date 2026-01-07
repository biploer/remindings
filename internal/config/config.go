package config

import "os"

type config struct {
	BotToken string
}

func MustConfig() config {
	botToken := os.Getenv("BOT_TOKEN")

	if botToken == "" {
		panic("BOT_TOKEN env variable is must be")
	}

	return config{
		BotToken: botToken,
	}
}
