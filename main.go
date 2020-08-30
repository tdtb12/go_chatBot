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

	log.Println("main " + constants.APIEnd)
}
