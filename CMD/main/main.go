package main

import (
	"fmt"
	"gocrudmongo/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.EmployeeRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running on: http://localhost:5005")
	log.Fatal(http.ListenAndServe(":5005", r))
}