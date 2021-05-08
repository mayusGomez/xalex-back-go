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

	log.Printf("request.Body: %+v", request.Body)
	quoteId := request.PathParameters["quoteId"]
	userId := request.QueryStringParameters["userId"]

	if len(request.Body) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Notes{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	var note billing.Notes
	err := json.Unmarshal([]byte(request.Body), &note)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Notes{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.QuoteMongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err = storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Notes{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	err = services.AddQuoteNotes(userId, quoteId, &note, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Notes{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(&note),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(note *billing.Notes) string {

	data, _ := json.Marshal(note)
	return string(data)

}
