package main

import (
	"chatBot/constants"
	"chatBot/router"
	"chatBot/utilities"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	log.Println("main " + constants.APIStart)
	fmt.Println("listen before")
	router.Listen()
	fmt.Println("listen after")
	utilities.ConnectMongo()

	log.Println("main " + constants.APIEnd)
}
