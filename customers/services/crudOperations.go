package services

import (
	"errors"
	"fmt"
	"log"
	"strings"

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

	customer.Name = strings.ToUpper(customer.Name)
	customer.LastName = strings.ToUpper(customer.LastName)
	customer.Email = strings.ToUpper(customer.Email)
	customer.IDType = strings.ToUpper(customer.IDType)
	customer.IDNumber = strings.ToUpper(customer.IDNumber)
	customer.Segment = strings.ToUpper(customer.Segment)

	err = s.Create(customer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int64, s storage.Storage) ([]customers.Customer, error) {

	users, err := s.GetByPage(IDUser, filterField, fielterData, pageNumber, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func Update(customer *customers.Customer, s storage.Storage) error {

	err := s.Update(customer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetCustomer(customerId string, s storage.Storage) (customers.Customer, error) {

	customer, err := s.Get(customerId)
	if err != nil {
		fmt.Println("Error, Receive err from storage", err)
		return customers.Customer{}, err
	}
	return customer, nil
}
