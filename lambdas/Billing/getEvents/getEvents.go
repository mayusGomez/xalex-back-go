package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/services"
	"github.com/mayusGomez/xalex/billing/storage/mongo"
)

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	IDUser := request.QueryStringParameters["userId"]
	filterField := request.QueryStringParameters["filterField"]
	filterData := request.QueryStringParameters["filterData"]
	pageNumberQ := request.QueryStringParameters["pageNumber"]
	pageSizeQ := request.QueryStringParameters["pageSize"]

	pageNumber, _ := strconv.ParseInt(pageNumberQ, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeQ, 10, 64)

	if IDUser == "" || pageNumber == 0 || pageSize == 0 {
		log.Println("Error, pageNumber or pageSize incorrect")
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	log.Println("IDUser:", IDUser)
	log.Println("filterField:", filterField)
	log.Println("filterData:", filterData)
	log.Println("pageNumberQ:", pageNumberQ)
	log.Println("pageSizeQ:", pageSizeQ)

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
			Body:       createResponse([]billing.Event{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	eventsData, err := services.GetEvensByPage(IDUser, filterField, filterData, pageNumber, pageSize, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createResponse([]billing.Event{}),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result len events:", len(eventsData))

	if len(eventsData) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse([]billing.Event{}),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(eventsData),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(events []billing.Event) string {

	data, _ := json.Marshal(events)
	return string(data)

}
