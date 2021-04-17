package billing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	InactiveServStatus ServiceStatus = 2
	ActiveServStatus   ServiceStatus = 1

	CancelAppointStatus   EventStatus = 1
	PendingAppointStatus  EventStatus = 2
	ExecutedAppointStatus EventStatus = 3

	AppointmentEvent EventType = 1
	BillingEvent     EventType = 2
)

type EventStatus int
type ServiceStatus int
type EventType int

type Service struct {
	IDmgo       primitive.ObjectID `bson:"_id" json:"-"`
	ID          string             `json:"id,omitempty" bson:"-"`
	IDUser      string             `json:"id_user,omitempty" bson:"id_user,omitempty"`
	Description string             `json:"description,omitempty"`
	Price       float32            `json:"price,omitempty" bson:"-"`
	Cost        float32            `json:"cost,omitempty" bson:"-"`
	PriceInt    int                `json:"-" bson:"price"`
	CostInt     int                `json:"-" bson:"cost"`
	Status      ServiceStatus      `json:"-"`
}

type DetailService struct {
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty" bson:"-"`
	Cost        float32 `json:"cost,omitempty" bson:"-"`
	PriceInt    int     `json:"-" bson:"price"`
	CostInt     int     `json:"-" bson:"cost"`
}

type EventCustomer struct {
	ID              string `json:"id,omitempty" bson:"-"`
	Name            string `json:"name,omitempty" bson:"name"`
	LastName        string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	MainMobilePhone string `json:"main_mobile_phone,omitempty" bson:"main_mobile_phone,omitempty"`
	Email           string `json:"email,omitempty" bson:"email,omitempty"`
	IDType          string `json:"id_type,omitempty" bson:"id_type,omitempty"`
	IDNumber        string `json:"id_number,omitempty" bson:"id_number,omitempty"`
}

type Event struct {
	IDmgo        primitive.ObjectID `bson:"_id" json:"-"`
	ID           string             `json:"id,omitempty" bson:"-"`
	IDUser       string             `json:"id_user,omitempty" bson:"id_user,omitempty"`
	Customer     EventCustomer      `json:"customer,omitempty"`
	EventType    EventType          `json:"event_type,omitempty"`
	Date         string             `json:"date,omitempty"`
	Datetime     time.Time          `json:"datetime,omitempty"`
	RegisterDate time.Time          `json:"register_date,omitempty"`
	Professional string             `json:"professional,omitempty"`
	Status       EventStatus        `json:"status,omitempty"`
	Note         string             `json:"note,omitempty"`
	Services     []DetailService    `json:"services,omitempty"`
}
