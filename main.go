package main

import (
	"fmt"

	"github.com/edniltonms17/api-go-rest/database"
	"github.com/edniltonms17/api-go-rest/routes"
)

const initialMesage string = "Starting our server Golang"

func main() {
	database.DataBaseConnection()
	fmt.Println(initialMesage)
	routes.HandleRequest()
}
