package router

import (
	"chatBot/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

func Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.RootController)

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
