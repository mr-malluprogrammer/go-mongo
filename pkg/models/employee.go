package models

import (
	"context"
	"gocrudmongo/pkg/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	EmployeeName string   `json:"empName"`
	Age          int      `json:"age"`
	Reporting    string   `json:"reporting"`
	EmployeeID   string   `json:"empID"`
	Project      *Project `json:"project"`
}

type Project struct {
	ProjectName     string `json:"projName"`
	ProjectLocation string `json:"projLoc"`
	ProjectID       int    `json:"projID"`
	ProjectStatus   string `json:"projStatus"`
}

var COLL = config.Connection().Database("youtube-employee").Collection("details")

func (e *Employee) CreateEmployeeDetail() *Employee {
	_, err := COLL.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func ShowAllEmployeeDetails() []Employee {
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Employees []Employee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		log.Fatal(err)
	}

	return Employees
}

func ShowEmployeeDetail(Id string) *Employee {
	var Employees Employee
	cursor := COLL.FindOne(context.TODO(), bson.M{"employeeid": Id})
	cursor.Decode(&Employees)
	return &Employees
}

func (e *Employee) UpdateEmployeeDetail(Id string) *Employee {
	var Employees Employee
	update := bson.M{
		"$set": e,
	}
	_,err := COLL.UpdateOne(context.TODO(), bson.M{"employeeid": Id}, update)
	if err != nil {
		panic(err)
	}
	cursor1 := COLL.FindOne(context.TODO(), bson.M{"employeeid": Id})
	cursor1.Decode(&Employees)
	return &Employees
}

func DeleteEmployeeDetail(Id string) []Employee {
	_, err := COLL.DeleteOne(context.TODO(), bson.M{"employeeid": Id})
	if err != nil {
		panic(err)
	}
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Employees []Employee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		log.Fatal(err)
	}

	return Employees
}