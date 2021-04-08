package storage

import (
	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/users/storage"
)

// Storage interface to Customer data access
type Storage interface {
	Connect() error
	Disconnect()
	Get(idUser string, idCustomer string) (customers.Customer, error)
	GetByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int, s storage.Storage) ([]customers.Customer, error)
	Create(customer *customers.Customer) error
}
