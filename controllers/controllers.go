package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edniltonms17/api-go-rest/models"
)

const homepage = "Home Page"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homepage)
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)

}
