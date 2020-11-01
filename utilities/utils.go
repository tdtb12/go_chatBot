package utilities

import (
	"chatBot/constants"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() {
	fmt.Println("connectMongo")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database(constants.DATABASE_BOT).Collection(constants.COLLECTION_LOG)

	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		fmt.Println("error la")
	}
	id := res.InsertedID
	fmt.Println(id)
}
