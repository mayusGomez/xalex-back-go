package mongo

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoStorage_Create(t *testing.T) {
	return
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
		IDUser:          "auth0|605fb236abbbeb006878e277",
		Name:            "Alexander",
		LastName:        "Gomez",
		MainMobilePhone: "3209876543",
		AuxMobilePhone:  auxMobPhone,
		Email:           "test@test.com",
		IDType:          "CC",
		IDNumber:        "888888888",
		Segment:         "Other",
		Location: &customers.Location{
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

func TestMongoStorage_GetByPage(t *testing.T) {
	return
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
		StringConn string
		Client     *mongo.Client
		Context    context.Context
	}
	type args struct {
		IDUser        string
		filterField   string
		filterPattern string
		pageNumber    int64
		pageSize      int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []customers.Customer
		wantErr bool
	}{
		{
			name: "TestMongoStorage_GetByPage ok by name",
			fields: fields{
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				IDUser:        "auth0|605fb236abbbeb006878e277",
				filterField:   "name",
				filterPattern: "ALEXANDER",
				pageNumber:    1,
				pageSize:      5,
			},
			want: []customers.Customer{
				customers.Customer{
					Name: "ALEXANDER",
				},
			},
			wantErr: false,
		},
		{
			name: "TestMongoStorage_GetByPage get all page 2",
			fields: fields{
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				IDUser:     "auth0|605fb236abbbeb006878e277",
				pageNumber: 2,
				pageSize:   3,
			},
			want: []customers.Customer{
				customers.Customer{
					Name: "CARLOS",
				},
			},
			wantErr: false,
		},
		{
			name: "TestMongoStorage_GetByPage get all empty",
			fields: fields{
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				IDUser:     "auth0|605fb236abbbeb006878e277",
				pageNumber: 10,
				pageSize:   3,
			},
			want: []customers.Customer{
				customers.Customer{
					Name: "CARLOS",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MongoStorage{
				StringConn: tt.fields.StringConn,
				Client:     tt.fields.Client,
				Context:    tt.fields.Context,
			}
			got, err := s.GetByPage(tt.args.IDUser, tt.args.filterField, tt.args.filterPattern, tt.args.pageNumber, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoStorage.GetByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) > 0 && got[0].Name != tt.want[0].Name {
				t.Errorf("MongoStorage.GetByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoStorage_Get(t *testing.T) {
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

	objId, _ := primitive.ObjectIDFromHex("606e85a8d9c1e3d55dde8680")

	type fields struct {
		StringConn string
		Client     *mongo.Client
		Context    context.Context
	}
	type args struct {
		idUser     string
		idCustomer string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    customers.Customer
		wantErr bool
	}{
		{
			name: "TestMongoStorage_Get_OK",
			fields: fields{
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				idUser:     "auth0|605fb236abbbeb006878e277",
				idCustomer: "606e85a8d9c1e3d55dde8680",
			},
			want: customers.Customer{
				IDmgo: objId, Name: "PATRICIA", LastName: "GOMEZ",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MongoStorage{
				StringConn: tt.fields.StringConn,
				Client:     tt.fields.Client,
				Context:    tt.fields.Context,
			}
			got, err := s.Get(tt.args.idCustomer)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoStorage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name == tt.want.Name && got.IDmgo != tt.want.IDmgo {
				t.Errorf("MongoStorage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
