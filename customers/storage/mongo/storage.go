package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (s *MongoStorage) Get(idCustomer string) (customers.Customer, error) {
	var customer customers.Customer
	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.CustommerCollection)

	objId, err := primitive.ObjectIDFromHex(idCustomer)
	if err != nil {
		log.Fatal(err)
		return customers.Customer{}, err
	}

	filter := bson.M{"_id": objId}
	err = coll.FindOne(s.Context, filter).Decode(&customer)
	if err != nil {
		log.Fatal(err)
		return customers.Customer{}, err
	}

	return customer, nil
}

// GetByPage customers of an User
func (s *MongoStorage) GetByPage(IDUser, filterField, filterPattern string, pageNumber, pageSize int64) ([]customers.Customer, error) {

	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.CustommerCollection)

	skips := pageSize * (pageNumber - 1)

	var customersList []customers.Customer
	findOpts := options.Find()
	findOpts.SetSkip(skips)
	findOpts.SetLimit(pageSize)
	findOpts.SetSort(bson.D{{"name", 1}})

	filter := bson.D{{"id_user", IDUser}}
	if filterField != "" && filterPattern != "" {
		filter = append(filter, bson.E{filterField, primitive.Regex{Pattern: filterPattern, Options: ""}})
	}

	cur, err := coll.Find(s.Context, filter, findOpts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(s.Context) {
		var customer customers.Customer
		err = cur.Decode(&customer)
		if err != nil {
			log.Fatal(err)
			continue
		}
		customersList = append(customersList, customer)
	}

	return customersList, nil
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

func (s *MongoStorage) Update(customer *customers.Customer) error {
	return nil
}
