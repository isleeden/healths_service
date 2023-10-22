package services

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

func GetHealthServices(servicesProps []entities.ServiceProps) []entities.HealthService {
	var stats []entities.HealthService
	ch := make(chan entities.HealthService, chBuffer)

	go checkServices(ch, servicesProps)

	for data := range ch {
		stats = append(stats, data)
	}

	return stats
}

func checkServices(ch chan entities.HealthService, servicesProps []entities.ServiceProps) {
	for _, service := range servicesProps {
		status := getServiceStatus(service.Url)

		ch <- entities.HealthService{
			Status: status,
			Name:   service.Name,
			Url:    service.Url,
		}
	}
	close(ch)
}

func getServiceStatus(url string) string {
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
