// Author: Amar
package main

import (
	"fmt"
)

type Address struct {
	Street string

	City string

	State string

	ZipCode int
}

type Student struct {
	Studentid int

	Name string

	Address
}

type Teacher struct {
	EmployeeID int

	Name string

	Salary float64

	Address
}

func defaultStudent(name string) *Student {
	var s Student
	s.Name = name
	s.Studentid = 12345
	s.Street = "Default St"
	s.City = "Default City"
	s.State = "Defaultville"
	s.ZipCode = 11121
	return &s
}

func defaultTeacher(name string) *Teacher {
	var t Teacher
	t.Name = name
	t.EmployeeID = 12345
	t.Salary = 13.50
	t.Street = "Default St"
	t.City = "Default City"
	t.State = "Defaultville"
	t.ZipCode = 11111
	return &t
}

func (s *Student) printInfo() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Student ID:", s.Studentid)
	fmt.Printf("Address: %s, %s, %s %d\n\n", s.Street, s.City, s.State, s.ZipCode)
}

func (t *Teacher) printInfo() {
	fmt.Println("Name:", t.Name)
	fmt.Println("Employee ID:", t.EmployeeID)
	fmt.Printf("Salary: $%0.2f\n", t.Salary)
	fmt.Printf("Address: %s, %s, %s %d\n\n", t.Street, t.City, t.State, t.ZipCode)
}

func main() {
	fmt.Println("A default student: ")
	s1 := defaultStudent("John")
	s1.printInfo()

	fmt.Println("A well off student: ")
	s2 := Student{Studentid: 34985, Name: "Martin", Address: Address{Street: "23922 E Dove Ln", City: "Chandler", State: "AZ", ZipCode: 85225}}
	s2.printInfo()

	fmt.Println("A default teacher: ")
	t1 := defaultTeacher("Mr. Smith")
	t1.printInfo()

	fmt.Println("A well off teacher: ")
	t2 := Teacher{EmployeeID: 73799, Name: "Dr. Richman", Salary: 34.50, Address: Address{Street: "77672 Freemond Dr", City: "Gilbert", State: "AZ", ZipCode: 85295}}
	t2.printInfo()
}
