package usecases

import (
	"fmt"

	"github.com/Shemistan/healths_service/internal/entities"
	"github.com/Shemistan/healths_service/internal/services"
)

func CheckServicesUsecase(serviceUrls []entities.ServiceProps) []entities.ServiceStatus {
	result := services.CheckServices(serviceUrls)

	fmt.Println("Monitoring Services...")

	return result
}
