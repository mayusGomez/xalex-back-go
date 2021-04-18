package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
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
	}{
		{
			name: "TestLambdaHandler001",
			args: args{
				request: events.APIGatewayProxyRequest{},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: false,
		},
	}
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
