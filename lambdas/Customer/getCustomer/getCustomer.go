package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/customers/services"
	"github.com/mayusGomez/xalex/customers/storage/mongo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	customerId := request.PathParameters["customerId"]
	log.Printf("request.customerId: %+v", customerId)
	userId := request.QueryStringParameters["userId"]
	log.Printf("request.userId: %+v", userId)

	if customerId == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&customers.Customer{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.MongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err := storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&customers.Customer{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	customer, err := services.GetCustomer(userId, customerId, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&customers.Customer{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result customer:", customer)

	if customer.ID == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&customer),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(&customer),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(customer *customers.Customer) string {

	data, _ := json.Marshal(customer)
	return string(data)

}
