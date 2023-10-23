package storage

import "github.com/Shemistan/healths_service/internal/entities"

type HealthProps interface {
	GetAll() []entities.HealthProps
}
