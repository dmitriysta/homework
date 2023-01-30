package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

var employees []Employee
var lastID int

func addEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastID++
	newEmployee.Id = lastID
	employees = append(employees, newEmployee)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Employee with ID %d has been added", newEmployee.Id)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, emp := range employees {
		if emp.Id == id {
			json, err := json.Marshal(emp)
			if err != nil {
				http.Error(w, "Could not encode employee to JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func dismissEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, emp := range employees {
		if emp.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func changeSalary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var emp Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.Id == id {
			employees[i].Salary = emp.Salary
			json.NewEncoder(w).Encode(employees[i])
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func changeDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var emp Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.Id == id {
			employees[i].Department = emp.Department
			json.NewEncoder(w).Encode(employees[i])
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	port := ":80"
	log.Println("Server listen on port:", port)
	http.HandleFunc("/add", addEmployee)
	http.HandleFunc("/employee", getEmployeeByID)
	http.HandleFunc("/dismiss", dismissEmployee)
	http.HandleFunc("/change-salary", changeSalary)
	http.HandleFunc("/employee", changeDepartment)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server: Could not listen and serve", err)
	}
}
