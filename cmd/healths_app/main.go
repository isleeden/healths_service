package main

import (
	"time"

	"github.com/Shemistan/healths_service/configs"
	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/presenters"
	"github.com/Shemistan/healths_service/internal/usecases"
)

func main() {
	config := configs.New()

	healthScheduler(config.ServiceUrls, config.Delay)
}

func healthScheduler(serviceUrls []entities.ServiceProps, delay time.Duration) {
	for {
		result := usecases.CheckServicesUsecase(serviceUrls)
		presenters.PrintHealthTable(result)
		time.Sleep(delay)
	}
}
