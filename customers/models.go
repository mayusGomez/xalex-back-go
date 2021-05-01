package customers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`
	Detail      string    `json:"detail,omitempty"`
}

type Date struct {
	Date  int `json:"date,omitempty"`
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type AuxMobilePhone struct {
	Number string `json:"number,omitempty"`
	Label  string `json:"label,omitempty"`
	Source string `json:"source,omitempty"`
}

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
}

// Customer Model of data
type Customer struct {
	IDmgo           primitive.ObjectID `bson:"_id" json:"-"`
	ID              string             `json:"id,omitempty" bson:"-"`
	IDUser          string             `json:"id_user,omitempty" bson:"id_user,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name"`
	LastName        string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	MainMobilePhone string             `json:"main_mobile_phone,omitempty" bson:"main_mobile_phone,omitempty"`
	AuxMobilePhone  []AuxMobilePhone   `json:"aux_mobile_phone,omitempty" bson:"aux_mobile_phone,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	IDType          string             `json:"id_type,omitempty" bson:"id_type,omitempty"`
	IDNumber        string             `json:"id_number,omitempty" bson:"id_number,omitempty"`
	Segment         string             `json:"segment,omitempty" bson:"segment,omitempty"`
	Location        *Location          `json:"location,omitempty" bson:"location,omitempty"`
	BirthDate       *Date              `json:"birth_date,omitempty" bson:"birth_date,omitempty"`
	Notes           []Notes            `json:"notes,omitempty" bson:"notes,omitempty"`
}
