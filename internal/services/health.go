package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Shemistan/healths_service/internal/entities"
)

const (
	goodStatus = "Working"
	badStatus  = "Not working"
)

func CheckServices(servicesProps []entities.ServiceProps) []entities.ServiceStatus {
	var stats []entities.ServiceStatus
	var wg sync.WaitGroup
	ch := make(chan entities.ServiceStatus)

	for _, service := range servicesProps {
		wg.Add(1)
		go checkService(&wg, ch, service)
	}

	wg.Wait()
	close(ch)

	for data := range ch {
		stats = append(stats, data)
	}

	return stats
}

func checkService(
	wg *sync.WaitGroup,
	ch chan<- entities.ServiceStatus,
	service entities.ServiceProps,
) {
	defer wg.Done()
	status := getServiceStatus(service.Url)

	ch <- entities.ServiceStatus{
		Status: status,
		Name:   service.Name,
		Url:    service.Url,
	}
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
