package configs

import (
	"time"

	"github.com/Shemistan/healths_service/internal/entities"
)

type Config struct {
	Delay       time.Duration
	ServiceUrls []entities.ServiceProps
}

func New() Config {
	return Config{
		Delay: 5 * time.Second,
		ServiceUrls: []entities.ServiceProps{
			{
				Url:  "https://google.com",
				Name: "Google",
			},
			{
				Url:  "https://payme.uz",
				Name: "Payme",
			},
		},
	}
}
