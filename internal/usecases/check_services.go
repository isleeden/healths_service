package usecases

import (
	"fmt"

	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/services"
)

func CheckServicesUsecase(serviceUrls []entities.ServiceProps) []entities.HealthService {
	result := services.GetHealthServices(serviceUrls)

	fmt.Println("Monitoring Services...")

	return result
}
