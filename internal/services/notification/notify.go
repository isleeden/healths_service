package notification

type NotificationService interface {
	SendError(message string)
}
