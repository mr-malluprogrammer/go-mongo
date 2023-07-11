package routes

import (
	"gocrudmongo/pkg/controller"

	"github.com/gorilla/mux"
)

var EmployeeRoutes = func(r *mux.Router) {
	r.HandleFunc("/employees", controller.ShowAllEmployeeDetails).Methods("GET")
	r.HandleFunc("/employee/{id}", controller.ShowEmployeeDetail).Methods("GET")
	r.HandleFunc("/employee", controller.CreateEmployeeDetail).Methods("POST")
	r.HandleFunc("/employee/{id}", controller.UpdateEmployeeDetail).Methods("PUT")
	r.HandleFunc("/employee/{id}", controller.DeleteEmployeeDetail).Methods("DELETE")
}
