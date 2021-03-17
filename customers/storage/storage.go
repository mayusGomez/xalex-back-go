package storage

import (
	"github.com/mayusGomez/xalex/customers"
)

// Storage interface to Customer data access
type Storage interface {
	Connect() error
	Disconnect()
	Get(idUser string, idCustomer string) (customers.Customer, error)
	GetByPage(idUser string, page int, size int) ([]customers.Customer, error)
	Create(customer *customers.Customer) error
}
