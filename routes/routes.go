package routes

import (
	"log"
	"net/http"

	controller "github.com/edniltonms17/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api/personalities", controller.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
