package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// method set id
func (e *Employee) SetId(id int) {
	e.id = id
}

// method set name
func (e *Employee) SetName(name string) {
	e.name = name
}

// method get id
func (e *Employee) GetId() int {
	return e.id
}

// method get name
func (e *Employee) GetName() string {
	return e.name
}

func main() {
	employee := Employee{id: 1, name: "Camilo"}
	fmt.Println(employee)
	employee.SetId(2)
	fmt.Println(employee)
	employee.SetName("Roosvell")
	fmt.Println(employee)
	fmt.Printf("mr. %v has %v id", employee.GetName(), employee.GetId())

}
