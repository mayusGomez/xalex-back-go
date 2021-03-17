package mongo

import (
	"context"
	"fmt"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	StringConn string
	Client     *mongo.Client
	Context    context.Context
}

func (s *MongoStorage) Connect() error {
	var err error

	s.Client, err = shared.GetMongoConn(s.StringConn, s.Context)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoStorage) Disconnect() {
	if s.Client != nil {
		s.Client.Disconnect(s.Context)
	}
}

// Get customer data by ID
func (s *MongoStorage) Get(idUser string, idCustomer string) (customers.Customer, error) {
	var customer customers.Customer

	return customer, nil
}

// GetByPage customers of an User
func (s *MongoStorage) GetByPage(idUser string, page int, size int) ([]customers.Customer, error) {

	var customer []customers.Customer
	return customer, nil
}

// Create customer data
func (s *MongoStorage) Create(customer *customers.Customer) error {

	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.CustommerCollection)

	customer.IDmgo = primitive.NewObjectID()
	result, err := coll.InsertOne(s.Context, customer)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	customer.ID = objectID.Hex()

	return err
}
