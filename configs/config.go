package configs

import (
	"time"

	"github.com/Shemistan/healths_service/internal/entities"
)

type Config struct {
	Delay          time.Duration
	ServiceToCheck []entities.HealthProps
}

func New() Config {
	return Config{
		Delay: 5 * time.Second,
		ServiceToCheck: []entities.HealthProps{
			{
				Endpoint: "https://google.com",
				Name:     "Google",
			},
			{
				Endpoint: "https://payme.uz/api",
				Name:     "Payme",
			},
		},
	}
}
