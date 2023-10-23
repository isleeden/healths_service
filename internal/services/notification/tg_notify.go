package notification

import (
	"context"

	"github.com/Shemistan/healths_service/configs"
	"github.com/nikoksr/notify/service/telegram"
)

const (
	errPrefix = "ERROR: "
)

type TgNotification struct {
	tgService *telegram.Telegram
	isEnabled bool
}

func New() NotificationService {
	cfg := configs.NewBotConfig()

	isEnabled := true

	tgService, err := telegram.New(cfg.TgBotToken)
	tgService.AddReceivers(824850459)

	if err != nil {
		isEnabled = false
	}

	return TgNotification{
		tgService: tgService,
		isEnabled: isEnabled,
	}
}

func (tgN TgNotification) SendError(message string) {
	tgN.tgService.Send(context.Background(), errPrefix, message)
}
