package services

import (
	"fmt"
	"strings"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/storage"
)

func CreateEvent(event *billing.Event, s storage.EventStorage) error {

	event.Professional = strings.ToUpper(event.Professional)
	event.Status = billing.PendingAppointStatus

	err := s.Create(event)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
