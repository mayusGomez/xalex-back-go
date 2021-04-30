package storage

import (
	"github.com/mayusGomez/xalex/customers"
)

// Storage interface to Customer data access
type Storage interface {
	Connect() error
	Disconnect()
	Get(idCustomer string) (customers.Customer, error)
	GetByPage(IDUser, filterField, filterPattern string, pageNumber, pageSize int64) ([]customers.Customer, int64, error)
	Create(customer *customers.Customer) error
	Update(customer *customers.Customer) error
}
