package controller

import (
	"fmt"
	"net/http"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	fmt.Print("root controller")
	w.Write([]byte("Welcome"))
}
