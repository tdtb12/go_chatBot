package main

import (
	"chatBot/router"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	log.Println("main start")
	router.Listen()
	log.Println("main end")
}
