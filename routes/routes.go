package routes

import (
	"log"
	"net/http"

	controller "github.com/edniltonms17/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	route := mux.NewRouter()
	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/api/personalities", controller.AllPersonalities).Methods("Get")
	route.HandleFunc("/api/personalities/{id}", controller.PersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", route))
}
