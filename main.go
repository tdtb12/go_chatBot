package main

import (
	"chatBot/constants"
	"chatBot/router"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	log.Println("main " + constants.APIStart)

	router.Listen()
	// connection := util.GetMongoConnection()
	// logCollection := connection.Database(constants.DATABASE_BOT).Collection(constants.COLLECTION_LOG)

	log.Println("main " + constants.APIEnd)
}
