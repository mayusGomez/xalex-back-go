package services

import (
	"fmt"
	"strings"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/storage"
)

func CreateService(service *billing.Service, s storage.ServiceStorage) error {

	service.Description = strings.ToUpper(service.Description)

	err := s.Create(service)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int64, s storage.ServiceStorage) ([]billing.Service, error) {

	services, err := s.GetByPage(IDUser, filterField, fielterData, pageNumber, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return services, nil
}

func Update(service *billing.Service, s storage.ServiceStorage) error {

	err := s.Update(service)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetService(serviceId string, s storage.ServiceStorage) (billing.Service, error) {

	service, err := s.Get(serviceId)
	if err != nil {
		fmt.Println("Error, Receive err from storage", err)
		return billing.Service{}, err
	}
	return service, nil
}
