package configs

import (
	"time"

	"github.com/Shemistan/healths_service/internal/entities"
)

type Config struct {
	Delay          time.Duration
	ServiceToCheck []entities.Service
}

func New() Config {
	return Config{
		Delay: 5 * time.Second,
		ServiceToCheck: []entities.Service{
			{
				Id:       1,
				Endpoint: "https://google.com",
				Name:     "Google",
			},
			{
				Id:       2,
				Endpoint: "https://payme.u/api",
				Name:     "Payme",
			},
			{
				Id:       3,
				Endpoint: "https://facebook.com",
				Name:     "Facebook",
			},
		},
	}
}
