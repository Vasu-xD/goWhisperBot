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




serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
clientOptions := options.Client().
    ApplyURI("mongodb+srv://vasu:vasu@cluster0.3gnlkhi.mongodb.net/?retryWrites=true&w=majority").
    SetServerAPIOptions(serverAPIOptions)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}


func GetDatabase() *mongo.Database {
	return GetClient().Database("whisper_bot")
}
