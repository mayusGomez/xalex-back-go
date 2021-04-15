package billing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CancelAppointStatus   = 0
	PendingAppointStatus  = 1
	ExecutedAppointStatus = 2

	AppointmentEvent = 0
	BillingEvent     = 1
)

type Service struct {
	IDmgo       primitive.ObjectID `bson:"_id" json:"-"`
	ID          string             `json:"id,omitempty" bson:"-"`
	IDUser      string             `json:"id_user,omitempty" bson:"id_user,omitempty"`
	Description string             `json:"description,omitempty"`
	Price       int                `json:"price,omitempty"`
	Cost        int                `json:"cost,omitempty"`
	Status      bool               `json:"status,omitempty"`
}

type DetailService struct {
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
	Cost        int    `json:"cost,omitempty"`
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
	EventType    int                `json:"event_type,omitempty"`
	Datetime     time.Time          `json:"datetime,omitempty"`
	RegisterDate time.Time          `json:"register_date,omitempty"`
	Professional string             `json:"professional,omitempty"`
	Status       string             `json:"status,omitempty"`
	Note         string             `json:"note,omitempty"`
	Services     []DetailService    `json:"services,omitempty"`
}