package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//1
	e := Employee{}
	fmt.Println(e)
	// 2
	employee := Employee{
		id:       1,
		name:     "Camilo",
		vacation: true,
	}
	fmt.Println(employee)
	// 3 returns an arrow value
	e3 := new(Employee)
	fmt.Println(*e3)
	e3.id = 5
	e3.name = "Roosvell"
	fmt.Println(*e3)
	// 4 create a constructor
	e4 := NewEmployee(1, "Andrew", true)
	fmt.Println(*e4)
}
