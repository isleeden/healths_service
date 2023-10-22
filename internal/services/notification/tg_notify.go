package notification

type TgNotification struct {
}

func New() NotificationService {
	return TgNotification{}
}
