// package utilities

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ConnectMongo() *mongo.Client {
// 	fmt.Println("connectMongo")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// collection := client.Database(constants.DATABASE_BOT).Collection(constants.COLLECTION_LOG)

// 	if err != nil {
// 		fmt.Println("error " + err)
// 	}

// 	return client
// }
