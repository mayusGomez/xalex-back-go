package mongo

import (
	"context"
	"errors"
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

	log.Println("Transform ObjectIDFromHex")
	objId, err := primitive.ObjectIDFromHex(idCustomer)
	if err != nil {
		log.Println("Error, get ObjectIDFromHex", err)
		return customers.Customer{}, err
	}

	filter := bson.M{"_id": objId}
	log.Println("Find One Document")
	err = coll.FindOne(s.Context, filter).Decode(&customer)
	if err != nil {
		log.Println("Error, Return empty customer;", err)
		return customers.Customer{}, nil
	}

	if customer.IDUser == "" {
		log.Println("Error, No idUser in data returned")
		return customers.Customer{}, nil
	}

	customer.ID = customer.IDmgo.Hex()

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
		log.Println("Error, Find:", err)
		return []customers.Customer{}, nil
	}

	for cur.Next(s.Context) {
		var customer customers.Customer
		err = cur.Decode(&customer)
		if err != nil {
			log.Fatal(err)
			continue
		}
		customer.ID = customer.IDmgo.Hex()
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

	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.CustommerCollection)
	var err error

	if customer.ID == "" {
		log.Fatal("Error: when try to update Customer, no ID received")
		return errors.New("Error: when try to update Customer, no ID received")
	}

	customer.IDmgo, err = primitive.ObjectIDFromHex(customer.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": customer.IDmgo},
		bson.D{
			{"$set", bson.D{
				{"id_user", customer.IDUser},
				{"name", customer.Name},
				{"last_name", customer.LastName},
				{"main_mobile_phone", customer.MainMobilePhone},
				{"email", customer.Email},
				{"id_type", customer.IDType},
				{"id_number", customer.IDNumber},
				{"segment", customer.Segment},
				{"aux_mobile_phone", customer.AuxMobilePhone},
				{"location", customer.Location},
				{"birth_date", customer.BirthDate},
				{"notes", customer.Notes},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
