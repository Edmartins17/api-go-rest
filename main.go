package main

import (
	"fmt"

	"github.com/edniltonms17/api-go-rest/models"
	"github.com/edniltonms17/api-go-rest/routes"
)

const initialMesage string = "Starting our server Golang"

func main() {
	models.Personalities = []models.Personality{
		{Name: "Nome 1", History: "Historia 1"},
		{Name: "Nome 2", History: "Historia 2"},
	}

	fmt.Println(initialMesage)
	routes.HandleRequest()
}
