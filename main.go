package main

import (
	"fmt"
	"log"
	"net/http"
)

const initialMesage string = "Iniciando nosso servidor Golang"
const homepage string = "Home Page"

func main() {
	fmt.Println(initialMesage)
	handleRequest()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homepage)
}

func handleRequest() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
