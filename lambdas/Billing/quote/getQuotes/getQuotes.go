package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mayusGomez/xalex/shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/services"
	"github.com/mayusGomez/xalex/billing/storage/mongo"
)

type response struct {
	Data     []billing.Quote `json:"data,omitempty"`
	ErrorMsg string          `json:"error_msg,omitempty"`
	Paging   *shared.Paging  `json:"paging,omitempty"`
}

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Request:", request)

	IDUser := request.QueryStringParameters["userId"]
	filterField := request.QueryStringParameters["filterField"]
	filterData := request.QueryStringParameters["filterData"]
	pageNumberQ := request.QueryStringParameters["pageNumber"]
	pageSizeQ := request.QueryStringParameters["pageSize"]

	statusStrList := strings.Split(request.QueryStringParameters["status"], ",")
	statusList := make([]billing.QuoteStatus, 0)

	if len(statusStrList) > 0 {
		for _, statusQ := range statusStrList {
			status, err := strconv.Atoi(statusQ)
			if err != nil {
				continue
			}
			statusList = append(statusList, billing.QuoteStatus(status))
		}
	}

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
	log.Println("statusList:", statusList)

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
			Body:       createErrorResponse("Processing error"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer storage.Disconnect()

	quotesData, totalData, err := services.GetQuotesByPage(IDUser, statusList, filterField, filterData, pageNumber, pageSize, &storage)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       createErrorResponse("Processing error"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	log.Println("Result len services:", len(quotesData))
	// log.Println("Result servicesData:", servicesData)

	if len(quotesData) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       createResponse(quotesData, 0),
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       createResponse(quotesData, totalData),
		StatusCode: http.StatusOK,
	}, nil
}

func createResponse(quotesData []billing.Quote, total int64) string {
	log.Println("createResponse quotesData:", quotesData)
	log.Println("total quotesData:", total)

	res := &response{
		Data: quotesData,
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
