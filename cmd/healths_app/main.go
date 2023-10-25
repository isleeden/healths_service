package main

import (
	"log"
	"time"

	"github.com/Shemistan/healths_service/configs"
	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/presenters"
	"github.com/Shemistan/healths_service/internal/services/health"
	"github.com/Shemistan/healths_service/internal/services/notification"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	cfg := configs.New()

	healthScheduler(
		cfg.ServiceToCheck,
		cfg.Delay,
	)
}

func healthScheduler(
	serviceUrls []entities.Service,
	delay time.Duration,
) {
	notify := notification.New()
	service := health.New(notify)

	for {
		result := service.GetHealth(serviceUrls)
		presenters.PrintHealthTable(result)
		time.Sleep(delay)
	}
}
