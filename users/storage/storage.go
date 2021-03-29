package storage

import (
	"github.com/mayusGomez/xalex/users"
)

// Storage interface to Customer data access
type Storage interface {
	Connect() error
	Disconnect()
	Get(id string) (users.User, error)
	Create(user *users.User) error
}
