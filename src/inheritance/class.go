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

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Roosvell"
	ftEmployee.age = 14
	ftEmployee.id = 1234
	fmt.Println(ftEmployee)
	// isn't allowed to make polimorphism directly,
	// we should use intefaces
	//GetMessage(ftEmployee)
}
