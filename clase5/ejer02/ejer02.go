package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e)
}

func main() {
	per := Person{ID: 1, Name: "Carlos", DateOfBirth: "01/02/2003"}
	emp := Employee{ID: 1, Position: "CEO", Person: Person{ID: 2, Name: "Laura", DateOfBirth: "03/04/2005"}}
	emp.PrintEmployee()
	emp2 := Employee{ID: 2, Position: "CTO", Person: per}
	emp2.PrintEmployee()
}
