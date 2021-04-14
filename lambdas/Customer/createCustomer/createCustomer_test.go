package main

import (
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
			name: "Lambdatest CreateUser no Body",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: "",
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"code":"0003","description":"No data provided","data":{}}`,
			},
			wantErr: false,
		},
		{
			name: "Lambdatest CreateUser wrong structure",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: `{"name": 19999}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"code":"0002","description":"Eror, request with wrong structure","data":{}}`,
			},
			wantErr: false,
		},
		{
			name: "Lambdatest CreateUser OK",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: `{"name": "Alex", "last_name":"Gomez", "main_mobile_phone":"3215674523"}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"code":"0000","description":"Eror, request with wrong structure","data":{}}`,
			},
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
			if got.Body != tt.want.Body {
				t.Errorf("LambdaHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
