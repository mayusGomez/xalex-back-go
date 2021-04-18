package storage

import (
	"github.com/mayusGomez/xalex/billing"
)

// ServiceStorage interface
type ServiceStorage interface {
	Connect() error
	Disconnect()
	Get(idService string) (billing.Service, error)
	GetByPage(IDUser, filterField, filterPattern string, pageNumber, pageSize int64) ([]billing.Service, error)
	Create(service *billing.Service) error
	Update(service *billing.Service) error
}

// EventStorage interface
type EventStorage interface {
	Connect() error
	Disconnect()
	Create(event *billing.Event) error
}
