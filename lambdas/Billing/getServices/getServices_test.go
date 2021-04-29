package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mayusGomez/xalex/billing"
)

func TestLambdaHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LambdaHandler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("LambdaHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LambdaHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createResponse(t *testing.T) {

	serv := make([]billing.Service, 0)
	serv = append(serv, billing.Service{ID: "prueba 001"})
	serv = append(serv, billing.Service{ID: "prueba 002"})

	type args struct {
		servicesData []billing.Service
		total        int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createResponse(tt.args.servicesData, tt.args.total); got != tt.want {
				t.Errorf("createResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createErrorResponse(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_createErrorResponseOK001",
			args: args{
				msg: "Test 001",
			},
			want: `{"error_msg":"Test 001"}`,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createErrorResponse(tt.args.msg); got != tt.want {
				t.Errorf("createErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
