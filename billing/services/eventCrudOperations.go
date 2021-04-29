package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/storage"
)

func CreateEvent(event *billing.Event, s storage.EventStorage) error {

	log.Printf("Service CreateEvent, event: %+v \n", event)
	if event.Datetime == (time.Time{}) {
		log.Println("datime empty, ", event.Datetime)
		event.Datetime = time.Now()
	}
	event.Date = event.Datetime.Format("20060102")
	event.Professional = strings.ToUpper(event.Professional)
	event.Status = billing.PendingAppointStatus
	event.RegisterDate = time.Now()
	event.SetMoneyToInt()

	if event.Customer != nil {
		event.Customer.Name = strings.ToUpper(event.Customer.Name)
		event.Customer.LastName = strings.ToUpper(event.Customer.LastName)
		event.Customer.Email = strings.ToLower(event.Customer.Email)
	}

	log.Printf("event to storage: %+v \n", event)

	err := s.Create(event)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetEvent(userId, eventId string, s storage.EventStorage) (billing.Event, error) {

	event, err := s.Get(eventId)

	if event.IDUser != userId {
		fmt.Println("Error, user not related", err)
		return billing.Event{}, nil
	}

	event.SetMoneyToFloat()
	if err != nil {
		fmt.Println("Error, Receive err from storage", err)
		return billing.Event{}, err
	}
	return event, nil
}

func GetEvensByPage(IDUser, filterField, fielterData string, pageNumber, pageSize int64, s storage.EventStorage) ([]billing.Event, error) {

	events, err := s.GetEventsByPage(IDUser, filterField, fielterData, pageNumber, pageSize)
	for _, event := range events {
		event.SetMoneyToFloat()
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return events, nil
}

func UpdateEvent(event *billing.Event, s storage.EventStorage) error {

	event.Date = event.Datetime.Format("20060102")
	event.Professional = strings.ToUpper(event.Professional)
	event.SetMoneyToInt()

	if event.Customer != nil {
		event.Customer.Name = strings.ToUpper(event.Customer.Name)
		event.Customer.LastName = strings.ToUpper(event.Customer.LastName)
		event.Customer.Email = strings.ToLower(event.Customer.Email)
	}

	log.Printf("event, update event: %+v \n", event)

	err := s.Update(event)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func PatchEvent(event *billing.Event, s storage.EventStorage) error {

	if event.ID == "" || event.IDUser == "" {
		log.Println("Patch event must have ID and userId")
		return errors.New("Patch event must have ID and userId")
	}

	if event.Datetime != (time.Time{}) {
		event.Date = event.Datetime.Format("20060102")
	}
	if event.Professional != "" {
		event.Professional = strings.ToUpper(event.Professional)
	}
	if event.Services != nil && len(event.Services) > 0 {
		event.SetMoneyToInt()
	}

	if event.Customer != nil {
		event.Customer.Name = strings.ToUpper(event.Customer.Name)
		event.Customer.LastName = strings.ToUpper(event.Customer.LastName)
		event.Customer.Email = strings.ToLower(event.Customer.Email)
	}

	log.Printf("event, patch event: %+v \n", event)

	err := s.Patch(event)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
