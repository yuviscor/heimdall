package notifier

import (
	"log"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type notifierService interface {
	Notify(serviceErr *domain.ServiceError) error
}

type NotifierManager struct {
	services []notifierService
}

func NewNotifierManager() *NotifierManager {
	return &NotifierManager{
		services: make([]notifierService, 0),
	}
}

func (manager *NotifierManager) AddService(service notifierService) {
	manager.services = append(manager.services, service)
}

func (manager *NotifierManager) Notify(serviceErr *domain.ServiceError) error {
	for _, service := range manager.services {
		if err := service.Notify(serviceErr); err != nil {
			log.Printf("[ERROR]: error in notify manager: %v\n", err)
		}
	}

	return nil
}
