package main

import (
	"fmt"
	"strconv"
)

func main() {
	// variables
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)
	myValue, err := strconv.ParseInt("abc", 0, 64)
	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println(myValue)
	}
	// maps
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])
	// slices
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(value)
		fmt.Println(index)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(value)
		fmt.Println(index)
	}
}
