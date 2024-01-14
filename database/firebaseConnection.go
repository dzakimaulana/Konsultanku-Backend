package database

import (
	"context"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client
var StorageClient *storage.Client

func FirebaseConnection() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("database/firebase-sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	client2, err := app.Storage(ctx)
	if err != nil {
		panic(err)
	}

	AuthClient = client
	StorageClient = client2
}
