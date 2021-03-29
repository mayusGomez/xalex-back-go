package main

import (
	"context"
	"encoding/json"
	"log"
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

	log.Printf("request.Body: %+v", request.Body)

	if len(request.Body) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse("0003", "No data provided", nil),
			StatusCode: 200,
		}, nil
	}

	var customer customers.Customer
	err := json.Unmarshal([]byte(request.Body), &customer)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       createResponse("0002", "Eror, request with wrong structure", nil),
			StatusCode: 200,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.MongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err = storage.Connect()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       createResponse("0001", "Processing Error", nil),
			StatusCode: 200,
		}, nil
	}
	defer storage.Disconnect()

	err = services.CreateUser(&customer, &storage)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       createResponse("0001", "Processing Error", nil),
			StatusCode: 200,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse("0000", "OK", &customer),
		StatusCode: 200,
	}, nil
}

func createResponse(code string, descrip string, customer *customers.Customer) string {

	resp := Result{
		Code:        code,
		Description: descrip,
	}
	if customer != nil {
		resp.Data = *customer
	}

	data, _ := json.Marshal(resp)
	return string(data)

}
