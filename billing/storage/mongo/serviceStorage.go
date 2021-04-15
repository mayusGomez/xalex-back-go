package mongo

import (
	"context"
	"log"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/shared"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.ServiceCollection)

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
