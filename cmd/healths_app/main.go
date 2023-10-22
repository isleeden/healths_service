package main

import (
	"time"

	"github.com/Shemistan/healths_service/configs"
	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/presenters"
	"github.com/Shemistan/healths_service/internal/services/health"
)

func main() {
	config := configs.New()

	healthScheduler(config.ServiceToCheck, config.Delay)
}

func healthScheduler(serviceUrls []entities.HealthProps, delay time.Duration) {
	service := health.New()

	for {
		result := service.GetHealth(serviceUrls)
		presenters.PrintHealthTable(result)
		time.Sleep(delay)
	}
}
