package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/services"
	"github.com/mayusGomez/xalex/billing/storage/mongo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	eventId := request.PathParameters["eventId"]
	userId := request.QueryStringParameters["userId"]
	log.Printf("request.eventId: %+v", eventId)
	log.Printf("request.userId: %+v", userId)

	if eventId == "" || userId == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Event{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.EventMongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err := storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Event{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	event, err := services.GetEvent(userId, eventId, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Event{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result event:", event)

	if event.ID == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Event{}),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(&event),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(event *billing.Event) string {

	data, _ := json.Marshal(event)
	return string(data)

}
