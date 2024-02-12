package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edniltonms17/api-go-rest/database"
	"github.com/edniltonms17/api-go-rest/models"
	"github.com/gorilla/mux"
)

const homepage = "Home Page"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homepage)
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var listPersolities []models.Personality
	database.DB.Find(&listPersolities)
	json.NewEncoder(w).Encode(listPersolities)

}

func PersonalityById(w http.ResponseWriter, r *http.Request) {
	personalities := mux.Vars(r)
	id := personalities["id"]
	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
