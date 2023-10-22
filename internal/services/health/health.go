package health

import "github.com/Shemistan/healths_service/internal/entities"

type HealthService interface {
	GetHealth(servicesProps []entities.HealthProps) []entities.Health
}
