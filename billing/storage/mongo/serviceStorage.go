package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/shared"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceMongoStorage struct {
	StringConn string
	Client     *mongo.Client
	Context    context.Context
}

func (s *ServiceMongoStorage) Connect() error {
	var err error

	s.Client, err = shared.GetMongoConn(s.StringConn, s.Context)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceMongoStorage) Disconnect() {
	if s.Client != nil {
		s.Client.Disconnect(s.Context)
	}
}

func (s *ServiceMongoStorage) Get(idService string) (billing.Service, error) {

	var service billing.Service
	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.ServiceCollection)

	log.Println("Transform ObjectIDFromHex")
	objId, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		log.Println("Error, get ObjectIDFromHex", err)
		return billing.Service{}, err
	}

	filter := bson.M{"_id": objId}
	log.Println("Find One Document")
	err = coll.FindOne(s.Context, filter).Decode(&service)
	if err != nil {
		log.Println("Error, Return empty service;", err)
		return billing.Service{}, nil
	}

	if service.IDUser == "" {
		log.Println("Error, No idUser in data returned")
		return billing.Service{}, nil
	}

	service.ID = service.IDmgo.Hex()

	return service, nil
}

func (s *ServiceMongoStorage) GetByPage(IDUser, filterField, filterPattern string, pageNumber, pageSize int64) ([]billing.Service, int64, error) {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.ServiceCollection)

	skips := pageSize * (pageNumber)

	var services []billing.Service
	findOpts := options.Find()
	findOpts.SetSkip(skips)
	findOpts.SetLimit(pageSize)
	findOpts.SetSort(bson.D{{"description", 1}})

	filter := bson.D{{"id_user", IDUser}, {"status", billing.ActiveServStatus}}
	if filterField != "" && filterPattern != "" {
		filter = append(filter, bson.E{filterField, primitive.Regex{Pattern: filterPattern, Options: ""}})
	}

	cur, err := coll.Find(s.Context, filter, findOpts)
	if err != nil {
		log.Println("Error, Find:", err)
		return []billing.Service{}, 0, nil
	}

	for cur.Next(s.Context) {
		var service billing.Service
		err = cur.Decode(&service)
		if err != nil {
			log.Fatal(err)
			continue
		}
		service.ID = service.IDmgo.Hex()
		services = append(services, service)
	}

	itemCount, _ := coll.CountDocuments(s.Context, filter)

	return services, itemCount, nil
}

func (s *ServiceMongoStorage) Create(service *billing.Service) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.ServiceCollection)

	service.IDmgo = primitive.NewObjectID()
	result, err := coll.InsertOne(s.Context, service)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	service.ID = objectID.Hex()

	return err
}

func (s *ServiceMongoStorage) Update(service *billing.Service) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.ServiceCollection)
	var err error

	if service.ID == "" {
		log.Fatal("Error: when try to update Service, no ID received")
		return errors.New("Error: when try to update Service, no ID received")
	}

	service.IDmgo, err = primitive.ObjectIDFromHex(service.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": service.IDmgo, "id_user": service.IDUser},
		bson.D{
			{"$set", bson.D{
				{"description", service.Description},
				{"price", service.PriceInt},
				{"cost", service.CostInt},
				{"status", service.Status},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
