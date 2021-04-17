package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/storage"
)

func CreateService(service *billing.Service, s storage.ServiceStorage) error {

	service.Description = strings.ToUpper(service.Description)
	service.SetMoneyToInt()
	service.Status = billing.ActiveServStatus

	err := s.Create(service)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int64, s storage.ServiceStorage) ([]billing.Service, error) {

	services, err := s.GetByPage(IDUser, filterField, fielterData, pageNumber, pageSize)
	for _, service := range services {
		service.SetMoneyToFloat()
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return services, nil
}

func Update(service *billing.Service, s storage.ServiceStorage) error {

	service.Description = strings.ToUpper(service.Description)
	service.SetMoneyToInt()
	service.Status = billing.ActiveServStatus

	log.Println("service, update service:", service)

	err := s.Update(service)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetService(userId, serviceId string, s storage.ServiceStorage) (billing.Service, error) {

	service, err := s.Get(serviceId)

	if service.IDUser != userId {
		fmt.Println("Error, user not related", err)
		return billing.Service{}, nil
	}

	service.SetMoneyToFloat()
	if err != nil {
		fmt.Println("Error, Receive err from storage", err)
		return billing.Service{}, err
	}
	return service, nil
}
