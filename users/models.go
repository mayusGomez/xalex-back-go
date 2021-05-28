package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
}

type Professional struct {
	Code int    `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type User struct {
	IDmgo         primitive.ObjectID `bson:"_id" json:"-"`
	ID            string             `json:"id,omitempty" bson:"id"`
	Name          string             `json:"name,omitempty" bson:"name"`
	LastName      string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	MobPhone      string             `json:"mob_phone,omitempty" bson:"mob_phone,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email"`
	Segment       string             `json:"segment,omitempty" bson:"segment,omitempty"`
	IDType        string             `json:"id_type,omitempty" bson:"id_type,omitempty"`
	IDCode        string             `json:"id_code,omitempty" bson:"id_code,omitempty"`
	Location      *Location          `json:"location,omitempty" bson:"location,omitempty"`
	Professionals []Professional     `json:"professionals,omitempty" bson:"professionals,omitempty"`
}
