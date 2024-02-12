package main

import (
	"fmt"

	"github.com/edniltonms17/api-go-rest/database"
	"github.com/edniltonms17/api-go-rest/models"
	"github.com/edniltonms17/api-go-rest/routes"
)

const initialMesage string = "Starting our server Golang"

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", History: "Historia 1"},
		{Id: 2, Name: "Nome 2", History: "Historia 2"},
		{Id: 3, Name: "Nome 3", History: "Historia 3"},
	}

	database.DataBaseConnection()
	fmt.Println(initialMesage)
	routes.HandleRequest()
}
