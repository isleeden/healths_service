package storage

import "github.com/Shemistan/healths_service/internal/entities"

type ServiceCheck interface {
	SaveCheck(check entities.ServiceCheck) error
}
