package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	id int
}
type FullTimeEmployee struct {
	// anonymous fields
	Person
	Employee
}
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

//interface
type PrintInfo interface {
	getMessage() string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}
func (tempEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Time Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Roosvell"
	ftEmployee.age = 14
	ftEmployee.id = 1234
	fmt.Println(ftEmployee)
	tempEmployee := TemporaryEmployee{}
	tempEmployee.name = "Camilo"
	tempEmployee.age = 33
	tempEmployee.id = 4321
	fmt.Println(tempEmployee)
	// isn't allowed to make polimorphism directly,
	// we should use intefaces
	getMessage(ftEmployee)
	getMessage(tempEmployee)
}
