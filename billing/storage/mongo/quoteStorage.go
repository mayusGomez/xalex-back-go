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

type QuoteMongoStorage struct {
	StringConn string
	Client     *mongo.Client
	Context    context.Context
}

func (s *QuoteMongoStorage) Connect() error {
	var err error

	s.Client, err = shared.GetMongoConn(s.StringConn, s.Context)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuoteMongoStorage) Disconnect() {
	if s.Client != nil {
		s.Client.Disconnect(s.Context)
	}
}

func (s *QuoteMongoStorage) Create(quote *billing.Quote) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.QuoteCollection)

	quote.IDmgo = primitive.NewObjectID()
	result, err := coll.InsertOne(s.Context, quote)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	quote.ID = objectID.Hex()

	return err
}

func (s *QuoteMongoStorage) Update(quote *billing.Quote) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.QuoteCollection)
	var err error

	if quote.ID == "" {
		log.Fatal("Error: when try to update Quote, no ID received")
		return errors.New("Error: when try to update Quote, no ID received")
	}

	quote.IDmgo, err = primitive.ObjectIDFromHex(quote.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": quote.IDmgo, "id_user": quote.IDUser},
		bson.D{
			{"$set", bson.D{
				{"customer", quote.Customer},
				{"professional", quote.Professional},
				{"status", quote.Status},
				{"description", quote.Description},
				{"services", quote.Services},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *QuoteMongoStorage) AddNote(userId, id string, note *billing.Notes) error {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.QuoteCollection)
	var err error

	if id == "" {
		log.Fatal("Error: when try to update Quote-note, no ID received")
		return errors.New("Error: when try to update Quote-note, no ID received")
	}

	idMgo, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = coll.UpdateOne(
		s.Context,
		bson.M{"_id": idMgo, "id_user": userId},
		bson.M{"$push": bson.M{"notes": note}},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *QuoteMongoStorage) Get(id string) (billing.Quote, error) {

	var quote billing.Quote
	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.QuoteCollection)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error, get ObjectIDFromHex", err)
		return billing.Quote{}, err
	}

	filter := bson.M{"_id": objId}
	err = coll.FindOne(s.Context, filter).Decode(&quote)
	if err != nil {
		log.Println("Error, Return empty quote;", err)
		return billing.Quote{}, nil
	}

	if quote.IDUser == "" {
		log.Println("Error, No idUser in data returned")
		return billing.Quote{}, nil
	}

	quote.ID = quote.IDmgo.Hex()

	return quote, nil
}

func (s *QuoteMongoStorage) GetQuotesByPage(IDUser string, quoteStatus []billing.QuoteStatus, filterField, filterPattern string, pageNumber, pageSize int64) ([]billing.Quote, int64, error) {

	db := s.Client.Database(billing.DBName)
	coll := db.Collection(billing.QuoteCollection)

	skips := pageSize * (pageNumber)

	var quotes []billing.Quote
	findOpts := options.Find()
	findOpts.SetSkip(skips)
	findOpts.SetLimit(pageSize)
	findOpts.SetSort(bson.D{{"register_date", -1}})

	filter := bson.D{{"id_user", IDUser}}
	if filterField != "" && filterPattern != "" {
		filter = append(filter, bson.E{filterField, primitive.Regex{Pattern: filterPattern, Options: ""}})
	}
	if len(quoteStatus) > 0 {
		filter = append(filter, bson.E{"status", bson.M{"$in": quoteStatus}})
	}

	cur, err := coll.Find(s.Context, filter, findOpts)
	if err != nil {
		log.Println("Error, Find:", err)
		return []billing.Quote{}, 0, nil
	}

	for cur.Next(s.Context) {
		var quote billing.Quote
		err = cur.Decode(&quote)
		if err != nil {
			log.Fatal(err)
			continue
		}
		quote.ID = quote.IDmgo.Hex()
		quotes = append(quotes, quote)
	}

	itemCount, _ := coll.CountDocuments(s.Context, filter)

	return quotes, itemCount, nil
}
