package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mayusGomez/xalex/customers"
)

func main() {
	lambda.Start(LambdaHandler)
}

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("request: %+v", request)
	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func createResponse(customers []customers.Customer) string {

	data, _ := json.Marshal(customers)
	return string(data)

}
