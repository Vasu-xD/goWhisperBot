package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Ctx    context.Context
	err    error
)

func GetClient() *mongo.Client {
	if Client != nil {
		return Client
	}

	Client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URI")))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(Ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Client
}

func GetDatabase() *mongo.Database {
	return GetClient().Database("whisper_bot")
}
