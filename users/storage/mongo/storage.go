package mongo

import (
	"context"
	"fmt"

	"github.com/mayusGomez/xalex/shared"
	"github.com/mayusGomez/xalex/users"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *MongoStorage) Create(user *users.User) error {

	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.UserCollection)

	user.IDmgo = primitive.NewObjectID()
	result, err := coll.InsertOne(s.Context, user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	return err
}

// Get User data by ID
func (s *MongoStorage) Get(id string) (users.User, error) {

	user := users.User{}
	filter := bson.D{{"id", id}}
	db := s.Client.Database(shared.DBName)
	coll := db.Collection(shared.UserCollection)
	err := coll.FindOne(s.Context, filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return users.User{}, nil
	}

	return user, nil
}
