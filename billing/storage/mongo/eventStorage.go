package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/shared"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventMongoStorage struct {
	StringConn string
	Client     *mongo.Client
	Context    context.Context
}

func (s *EventMongoStorage) Connect() error {
	var err error

	s.Client, err = shared.GetMongoConn(s.StringConn, s.Context)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventMongoStorage) Disconnect() {
	if s.Client != nil {
		s.Client.Disconnect(s.Context)
	}
}

func (s *EventMongoStorage) Create(event *billing.Event) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.EventCollection)

	event.IDmgo = primitive.NewObjectID()
	result, err := coll.InsertOne(s.Context, event)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	event.ID = objectID.Hex()

	return err
}

func (s *EventMongoStorage) Update(event *billing.Event) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.EventCollection)
	var err error

	if event.ID == "" {
		log.Fatal("Error: when try to update Event, no ID received")
		return errors.New("Error: when try to update Event, no ID received")
	}

	event.IDmgo, err = primitive.ObjectIDFromHex(event.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": event.IDmgo},
		bson.D{
			{"$set", bson.D{
				{"id_user", event.IDUser},
				{"customer", event.Customer},
				{"event_type", event.EventType},
				{"date", event.Date},
				{"start_date", event.StartDate},
				{"professional", event.Professional},
				{"status", event.Status},
				{"note", event.Note},
				{"services", event.Services},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *EventMongoStorage) Patch(event *billing.Event) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.EventCollection)
	var err error

	if event.ID == "" || event.IDUser == "" {
		log.Fatal("Error: when try to patch Event, no ID or IDUser received")
		return errors.New("Error: when try to patch Event, no ID or IDUser received")
	}

	idObj, err := primitive.ObjectIDFromHex(event.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	updateBson := bson.D{}
	if event.Customer != nil {
		updateBson = append(updateBson, bson.E{"customer", event.Customer})
	}

	if event.EventType != 0 {
		updateBson = append(updateBson, bson.E{"event_type", event.EventType})
	}

	if event.Date != "" {
		updateBson = append(updateBson, bson.E{"date", event.Date})
	}

	if event.StartDate != (time.Time{}) {
		updateBson = append(updateBson, bson.E{"date_time", event.StartDate})
	}

	if event.Professional != "" {
		updateBson = append(updateBson, bson.E{"professional", event.Professional})
	}

	if event.Status != 0 {
		updateBson = append(updateBson, bson.E{"status", event.Status})
	}

	if event.Note != "" {
		updateBson = append(updateBson, bson.E{"note", event.Note})
	}

	if event.Services != nil || len(event.Services) != 0 {
		updateBson = append(updateBson, bson.E{"services", event.Services})
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": idObj, "id_user": event.IDUser},
		bson.D{
			{"$set", updateBson},
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *EventMongoStorage) Get(idEvent string) (billing.Event, error) {

	var event billing.Event
	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.EventCollection)

	log.Println("Transform ObjectIDFromHex")
	objId, err := primitive.ObjectIDFromHex(idEvent)
	if err != nil {
		log.Println("Error, get ObjectIDFromHex", err)
		return billing.Event{}, err
	}

	filter := bson.M{"_id": objId}
	log.Println("Find One Document")
	err = coll.FindOne(s.Context, filter).Decode(&event)
	if err != nil {
		log.Println("Error, Return empty event;", err)
		return billing.Event{}, nil
	}

	if event.IDUser == "" {
		log.Println("Error, No idUser in data returned")
		return billing.Event{}, nil
	}

	event.ID = event.IDmgo.Hex()

	return event, nil
}

func (s *EventMongoStorage) GetEventsByPage(IDUser, filterField, filterPattern string, pageNumber, pageSize int64) ([]billing.Event, error) {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.EventCollection)

	skips := pageSize * (pageNumber - 1)

	var events []billing.Event
	findOpts := options.Find()
	findOpts.SetSkip(skips)
	findOpts.SetLimit(pageSize)
	findOpts.SetSort(bson.D{{"datetime", 1}})

	filter := bson.D{{"id_user", IDUser}}
	if filterField != "" && filterPattern != "" {
		filter = append(filter, bson.E{filterField, primitive.Regex{Pattern: filterPattern, Options: ""}})
	}

	cur, err := coll.Find(s.Context, filter, findOpts)
	if err != nil {
		log.Println("Error, Find:", err)
		return []billing.Event{}, nil
	}

	for cur.Next(s.Context) {
		var event billing.Event
		err = cur.Decode(&event)
		if err != nil {
			log.Fatal(err)
			continue
		}
		event.ID = event.IDmgo.Hex()
		events = append(events, event)
	}

	return events, nil
}
