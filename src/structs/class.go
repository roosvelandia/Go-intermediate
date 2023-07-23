package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	employee := Employee{id: 1, name: "Camilo"}
	fmt.Println(employee)
	employee.id = 2
	fmt.Println(employee)
}
