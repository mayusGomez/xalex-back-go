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

	quoteId := request.PathParameters["quoteId"]
	userId := request.QueryStringParameters["userId"]
	log.Printf("request.quoteId: %+v", quoteId)

	if quoteId == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Quote{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.QuoteMongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err := storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Quote{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	quote, err := services.GetQuote(userId, quoteId, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Quote{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result Quote:", quote)

	if quote.ID == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&quote),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(&quote),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(quote *billing.Quote) string {

	data, _ := json.Marshal(quote)
	return string(data)

}
