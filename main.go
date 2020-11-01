package main

import (
	"chatBot/constants"
	"chatBot/router"
	"chatBot/utilities"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	log.Println("main " + constants.APIStart)

	router.Listen()
	utilities.ConnectMongo()

	log.Println("main " + constants.APIEnd)
}
