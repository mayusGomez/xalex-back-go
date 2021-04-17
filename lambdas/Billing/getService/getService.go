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

	serviceId := request.PathParameters["serviceId"]
	userId := request.QueryStringParameters["userId"]
	log.Printf("request.serviceId: %+v", serviceId)

	if serviceId == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storage := mongo.ServiceMongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err := storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	service, err := services.GetService(userId, serviceId, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result service:", service)

	if service.ID == "" {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&service),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(&service),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(service *billing.Service) string {

	data, _ := json.Marshal(service)
	return string(data)

}
