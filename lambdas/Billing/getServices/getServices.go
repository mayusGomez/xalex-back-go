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

	"github.com/mayusGomez/xalex/shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/services"
	"github.com/mayusGomez/xalex/billing/storage/mongo"
)

type response struct {
	Data     []billing.Service `json:"data,omitempty"`
	ErrorMsg string            `json:"error_msg,omitempty"`
	Paging   *shared.Paging    `json:"paging,omitempty"`
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

	if IDUser == "" || pageNumber < 0 || pageSize <= 0 {
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

	storage := mongo.ServiceMongoStorage{
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

	servicesData, totalData, err := services.GetByPage(IDUser, filterField, filterData, pageNumber, pageSize, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createErrorResponse("Processing error"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result len services:", len(servicesData))
	log.Println("Result servicesData:", servicesData)

	if len(servicesData) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(servicesData, 0),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(servicesData, totalData),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(servicesData []billing.Service, total int64) string {
	log.Println("createResponse servicesData:", servicesData)
	log.Println("total servicesData:", total)

	res := &response{
		Data: servicesData,
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
