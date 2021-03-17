package mongo

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/shared"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoStorage_Create(t *testing.T) {

	strConn := os.Getenv("STR_MONGO_CONN")
	log.Println("strConn:", strConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := shared.GetMongoConn(strConn, ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Disconnect(ctx)

	type fields struct {
		Env        string
		StringConn string
		Client     *mongo.Client
		Context    context.Context
	}
	type args struct {
		customer *customers.Customer
	}

	auxMobPhone := make([]customers.AuxMobilePhone, 0)
	auxMobilePhone := customers.AuxMobilePhone{
		Number: "3109991818",
		Label:  "Home",
		Source: "Manual",
	}
	auxMobPhone = append(auxMobPhone, auxMobilePhone)

	notes := make([]customers.Notes, 0)
	not01 := customers.Notes{
		CreatedAt:   time.Now(),
		Description: "Note 001",
		Detail:      "Note Detail",
	}
	notes = append(notes, not01)

	customer := customers.Customer{
		Name:            "Alexander",
		LastName:        "Gomez",
		MainMobilePhone: "3209876543",
		AuxMobilePhone:  auxMobPhone,
		Email:           "test@test.com",
		IDType:          "CC",
		IDNumber:        "888888888",
		Segment:         "Other",
		Location: customers.Location{
			Country: "COP",
			City:    "Bogota",
			Address: "Calle 45 # 56-44",
		},
		Notes: notes,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert CustomerDataMongo 001",
			fields: fields{
				Env:        "Dev",
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				customer: &customer,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MongoStorage{
				StringConn: tt.fields.StringConn,
				Client:     tt.fields.Client,
				Context:    tt.fields.Context,
			}
			if err := s.Create(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("MongoStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
