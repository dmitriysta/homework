package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	SecondName string `json:"secondname"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Salary     string `json:"salary"`
	Department string `json:"department"`
}

var employees = []Employee{}
var listToJSON []string

func (e *Employee) addEmployee() {
	//Add new employee
	employees = []Employee{
		Employee{
			Id:         1,
			Name:       "Kirill",
			Surname:    "Petrov",
			SecondName: "Ivanovich",
			Address:    "Moscow",
			Phone:      "9491543482",
			Salary:     "100000",
			Department: "IT",
		},
	}

	//Create JSON
	for i := 0; i < len(employees); i++ {
		temp, err := json.Marshal(employees[i])
		if err != nil {
			log.Fatal("error convert to json", err)
		}
		listToJSON = append(listToJSON, string(temp))
	}
}

func (e *Employee) dismissEmployee() []string {
	var amountID string
	fmt.Println("Enter employee's ID, which needs to be dismiss:")
	fmt.Scan(&amountID)
	for i := 0; i < len(employees); i++ {
		if string(listToJSON[i][6]) == amountID {
			return append(listToJSON[:i], listToJSON[i+1:]...)
		}
	}
	return listToJSON
}

func (e *Employee) changeSalary() {
	var amountID string
	var newSalary string
	fmt.Println("Enter employee's ID, which needs to be change salary:")
	fmt.Scan(&amountID)
	fmt.Println("Enter new salary for this employee:")
	fmt.Scan(&newSalary)
	for i := 0; i < len(employees); i++ {
		if string(listToJSON[i][6]) == amountID {
			ind := strings.Index(listToJSON[i], "Salary")
			listToJSON[i] = strings.Replace(listToJSON[i], listToJSON[i][ind+8:ind+14], newSalary, 1)

		}
	}
}

func (e *Employee) changeDepartment() {
	var amountID string
	var newDepartment string
	fmt.Println("Enter employee's ID, which needs to be change department:")
	fmt.Scan(&amountID)
	fmt.Println("Enter new department for this employee:")
	fmt.Scan(&newDepartment)
	for i := 0; i < len(employees); i++ {
		if string(listToJSON[i][6]) == amountID {
			ind := strings.Index(listToJSON[i], "Department")
			listToJSON[i] = strings.Replace(listToJSON[i], listToJSON[i][ind+13:ind+16], newDepartment, 1)

		}
	}
}

func (e *Employee) getInfoById() {
	var amountID string
	fmt.Println("Enter employee's ID:")
	fmt.Scan(&amountID)
	for i := 0; i < len(employees); i++ {
		if string(listToJSON[i][6]) == amountID {
			fmt.Println("All info about ID:", listToJSON[i])
		}
	}
}