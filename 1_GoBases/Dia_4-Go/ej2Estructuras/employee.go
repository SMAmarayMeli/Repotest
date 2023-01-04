package main

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Id       int
	Position string
	Per      Person
}

func (e Employee) printEmployee() {
	fmt.Printf("Nombre: %v \nID: %v \nPosition: %v \n", e.Per.Name, e.Per.Id, e.Position)
}

func main() {
	p := Person{42589075, "Pepe", "12/6/24"}
	e := Employee{50, "Software developer", p}

	e.printEmployee()
}
