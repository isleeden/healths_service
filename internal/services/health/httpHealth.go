package health

import (
	"fmt"
	"net/http"

	"github.com/Shemistan/healths_service/internal/entities"
)

const (
	goodStatus = "Working"
	badStatus  = "Not working"

	chBuffer = 5
)

type HttpHealth struct {
}

func New() HealthService {
	return HttpHealth{}
}

func (service HttpHealth) GetHealth(healthProps []entities.HealthProps) []entities.Health {
	var healths []entities.Health
	ch := make(chan entities.Health, chBuffer)

	go checkHealth(ch, healthProps)

	for data := range ch {
		healths = append(healths, data)
	}

	return healths
}

func checkHealth(ch chan entities.Health, healthProps []entities.HealthProps) {
	for _, service := range healthProps {
		status := getHealthStatus(service.Url)

		ch <- entities.Health{
			Status: status,
			Name:   service.Name,
			Url:    service.Url,
		}
	}
	close(ch)
}

func getHealthStatus(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return defineStatusMessage(false)
	}

	return defineStatusMessage(resp.StatusCode == 200)
}

func defineStatusMessage(isOk bool) string {
	var message string
	if isOk {
		message = goodStatus
	} else {
		message = badStatus
	}
	return message
}
