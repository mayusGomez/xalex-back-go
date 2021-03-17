package services

import (
	"fmt"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/customers/storage"
)

func CreateUser(customer *customers.Customer, s storage.Storage) error {

	customer.ID = ""

	err := s.Create(customer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
