package routes

import (
	"log"
	"net/http"

	controller "github.com/edniltonms17/api-go-rest/controllers"
	"github.com/edniltonms17/api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	route := mux.NewRouter()
	route.Use(middleware.ContentTypeMiddleware)
	route.HandleFunc("/", controller.Home)
	const personalitiesURL string = "/api/personalities"
	route.HandleFunc(personalitiesURL, controller.AllPersonalities).Methods("Get")
	route.HandleFunc(personalitiesURL+"/{id}", controller.PersonalityById).Methods("Get")
	route.HandleFunc(personalitiesURL, controller.CreatePersonality).Methods("Post")
	route.HandleFunc(personalitiesURL+"/{id}", controller.DeletePersonality).Methods("Delete")
	route.HandleFunc(personalitiesURL, controller.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(route)))
}
