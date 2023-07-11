package controller

import (
	"encoding/json"
	"gocrudmongo/pkg/models"
	"gocrudmongo/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowAllEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	showEmployee := models.ShowAllEmployeeDetails()
	res,_ := json.Marshal(showEmployee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowEmployeeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := models.ShowEmployeeDetail(empID)
	res,_ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEmployeeDetail(w http.ResponseWriter, r *http.Request) {
	CreateEmployee := &models.Employee{}
	utils.ParseBody(r, CreateEmployee)
	employee := CreateEmployee.CreateEmployeeDetail()
	res,_ := json.Marshal(employee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateEmployeeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	UpdateEmployee := &models.Employee{}
	utils.ParseBody(r, UpdateEmployee)
	empDetailstoUpdate := UpdateEmployee.UpdateEmployeeDetail(empID)
	res,_ := json.Marshal(empDetailstoUpdate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteEmployeeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := models.DeleteEmployeeDetail(empID)
	res,_ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
