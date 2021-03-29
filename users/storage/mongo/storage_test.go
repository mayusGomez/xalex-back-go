package mongo

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/mayusGomez/xalex/shared"
	"github.com/mayusGomez/xalex/users"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoStorage_Create(t *testing.T) {

	strConn := os.Getenv("STR_MONGO_CONN")
	log.Println("strConn:", strConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := shared.GetMongoConn(strConn, ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Disconnect(ctx)

	user := users.User{
		ID:       "auth0|605fb236abbbeb006878e277",
		Name:     "Alonso",
		LastName: "Higuita",
		Email:    "alexander.gomez.higuita@gmail.com",
	}

	type fields struct {
		StringConn string
		Client     *mongo.Client
		Context    context.Context
	}
	type args struct {
		user *users.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestMongoStorage_Create 01",
			fields: fields{
				StringConn: strConn,
				Client:     client,
				Context:    ctx,
			},
			args: args{
				user: &user,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MongoStorage{
				StringConn: tt.fields.StringConn,
				Client:     tt.fields.Client,
				Context:    tt.fields.Context,
			}
			if err := s.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("MongoStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
