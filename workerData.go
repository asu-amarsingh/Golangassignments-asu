package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

type employee struct {
	name   string
	age    int
	salary float64
}

// default employee
func defaultEmployee(name string) employee {
	var e employee
	e.name = name
	e.age = 18
	e.salary = 13.50

	return e
}

func main() {
	//create CSV file to store name, age, and salary
	empData, err := os.Create("employeedata.csv")
	if err != nil {
		panic(err)
	}
	// use defer to close csv file
	defer empData.Close()

	//call new csv writer
	writer := csv.NewWriter(empData)

	//create employees and write to the csv file
	e1 := employee{name: "Marty", age: 24, salary: 24.00}
	e2 := defaultEmployee("Katy")
	e3 := defaultEmployee("John")
	e4 := employee{name: "Arnold", age: 50, salary: 54.30}

	emps := []employee{e1, e2, e3, e4}

	for _, e := range emps {
		row := []string{e.name, fmt.Sprintf("%d", e.age), fmt.Sprintf("$%0.2f", e.salary)}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
	writer.Flush()

	//read data out of the file and display it *sorted* by name (use sort package)
	empData, err = os.Open("employeedata.csv")
	if err != nil {
		panic(err)
	}
	defer empData.Close()

	reader := csv.NewReader(empData)
	database, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	sort.Slice(database, func(i, j int) bool {
		return database[i][0] < database[j][0]
	})

	fmt.Printf("Printing Name | Age | Wage: \n")

	for _, data := range database {
		fmt.Println(data)
	}
}
