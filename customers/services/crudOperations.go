package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/customers/storage"
)

func validateCreateUser(customer *customers.Customer) (string, error) {
	if customer.ID != "" {
		return "Code invalid", errors.New("Customer with ID")
	}
	if len(customer.Name) == 0 {
		return "Name invalid", errors.New("Customer without Name")
	}

	return "", nil
}

func CreateUser(customer *customers.Customer, s storage.Storage) error {

	errDesc, err := validateCreateUser(customer)
	if err != nil {
		log.Println("Error validating data:", err)
		return errors.New(errDesc)
	}

	err = s.Create(customer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int, s storage.Storage) ([]customers.Customer, error) {

	users, err := s.GetByPage(IDUser, filterField, fielterData, pageNumber, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}
