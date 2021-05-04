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
	InteractionEvent EventType = 3

	CanceledQuote  QuoteStatus = 1
	StartedQuote   QuoteStatus = 2
	FinalizedQuote QuoteStatus = 3
	PendingQuote   QuoteStatus = 4
	DiscartedQuote QuoteStatus = 5
)

type EventStatus int
type ServiceStatus int
type EventType int
type QuoteStatus int

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
	Name            string `json:"name,omitempty" bson:"name,omitempty"`
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
	Customer     *EventCustomer     `json:"customer,omitempty" bson:"customer,omitempty"`
	EventType    EventType          `json:"event_type,omitempty" bson:"event_type"`
	Date         string             `json:"-" bson:"date,omitempty"`
	StartDate    time.Time          `json:"start_date,omitempty" bson:"date_time,omitempty"`
	EndDate      time.Time          `json:"end_date,omitempty"`
	RegisterDate time.Time          `json:"register_date,omitempty" bson:"register_date,omitempty"`
	Professional string             `json:"professional,omitempty" bson:"professional,omitempty"`
	Status       EventStatus        `json:"status,omitempty" bson:"status,omitempty"`
	Note         string             `json:"note,omitempty" bson:"note,omitempty"`
	Services     []DetailService    `json:"services,omitempty" bson:"services,omitempty"`
	Description  string             `json:"description,omitempty"`
}

type Notes struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Detail    string    `json:"detail,omitempty"`
}

type Quote struct {
	IDmgo        primitive.ObjectID `bson:"_id" json:"-"`
	ID           string             `json:"id,omitempty" bson:"-"`
	Code         string             `json:"code,omitempty"`
	RegisterDate time.Time          `json:"register_date,omitempty"`
	IDUser       string             `json:"id_user,omitempty" bson:"id_user,omitempty"`
	Customer     *EventCustomer     `json:"customer,omitempty" bson:"customer,omitempty"`
	Professional string             `json:"professional,omitempty" bson:"professional,omitempty"`
	Status       QuoteStatus        `json:"status,omitempty"`
	Description  string             `json:"description,omitempty"`
	Notes        []Notes            `json:"notes,omitempty" bson:"notes,omitempty"`
	Services     []DetailService    `json:"services,omitempty"`
}
