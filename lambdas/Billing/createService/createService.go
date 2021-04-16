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

	if len(request.Body) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	var service billing.Service
	err := json.Unmarshal([]byte(request.Body), &service)
	if err != nil {
		fmt.Println(err)
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
	err = storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	err = services.CreateService(&service, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse(&billing.Service{}),
			StatusCode: http.StatusInternalServerError,
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
