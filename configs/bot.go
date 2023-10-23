package configs

import "os"

const (
	botToken = "TG_BOT_TOKEN"
)

type BotConfig struct {
	TgBotToken string
}

func NewBotConfig() BotConfig {
	return BotConfig{
		TgBotToken: os.Getenv(botToken),
	}
}
