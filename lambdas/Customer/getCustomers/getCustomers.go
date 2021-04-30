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
	"github.com/mayusGomez/xalex/customers"
	"github.com/mayusGomez/xalex/customers/services"
	"github.com/mayusGomez/xalex/customers/storage/mongo"
	"github.com/mayusGomez/xalex/shared"
)

type response struct {
	Data     []customers.Customer `json:"data,omitempty"`
	ErrorMsg string               `json:"error_msg,omitempty"`
	Paging   *shared.Paging       `json:"paging,omitempty"`
}

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

	if IDUser == "" || pageNumber < 0 || pageSize == 0 {
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

	storage := mongo.MongoStorage{
		StringConn: os.Getenv("STR_MONGO_CONN"),
		Context:    ctx,
	}
	err := storage.Connect()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createErrorResponse("Processing error"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	customersData, totalData, err := services.GetByPage(IDUser, filterField, filterData, pageNumber, pageSize, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createErrorResponse("Processing error"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result len customers:", len(customersData))

	if len(customersData) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(customersData, 0),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(customersData, totalData),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(customersData []customers.Customer, total int64) string {

	res := &response{
		Data: customersData,
		Paging: &shared.Paging{
			Total: total,
		},
	}

	log.Printf("response: %+v \n", res)

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Error, marshal response:", err)
	}
	log.Println("data:", string(data))

	return string(data)

}

func createErrorResponse(msg string) string {

	res := &response{
		ErrorMsg: msg,
	}
	data, _ := json.Marshal(res)
	return string(data)

}
