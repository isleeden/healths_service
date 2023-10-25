package health

import (
	"fmt"
	"net/http"

	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/services/notification"
)

const (
	goodStatus = "Working"
	badStatus  = "Not working"

	chBuffer = 5
)

type HttpHealth struct {
	notify notification.NotificationService
}

func New(notify notification.NotificationService) HealthService {
	return HttpHealth{
		notify: notify,
	}
}

func (health HttpHealth) GetHealth(healthProps []entities.Service) []entities.Health {
	var healths []entities.Health
	ch := make(chan entities.Health, chBuffer)

	go health.checkHealth(ch, healthProps)

	for data := range ch {
		healths = append(healths, data)
	}

	return healths
}

func (health HttpHealth) checkHealth(ch chan entities.Health, healthProps []entities.Service) {
	for _, service := range healthProps {
		status := health.getHealthStatus(service.Endpoint)

		ch <- entities.Health{
			Status: status,
			Name:   service.Name,
			Url:    service.Endpoint,
		}
	}
	close(ch)
}

func (health HttpHealth) getHealthStatus(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: ", err)
		health.notify.SendError(err.Error())
		return health.defineStatusMessage(false)
	}

	return health.defineStatusMessage(resp.StatusCode == 200)
}

func (health HttpHealth) defineStatusMessage(isOk bool) string {
	var message string
	if isOk {
		message = goodStatus
	} else {
		message = badStatus
	}
	return message
}
