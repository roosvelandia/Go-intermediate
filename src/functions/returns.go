package main

import "fmt"

// function rigid
func sum(a, b int) int {
	return a + b
}

// variadic function turns imput into a slice with return
func sumVariable(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// variadic function turns imput into a slice without return
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// variadic function return values without name
func getValues(x int) (int, int, int) {
	return 1 * x, 3 * x, 4 * x
}

// variadic function return values with name
func getValuesWithName(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sumVariable(1))
	fmt.Println(sumVariable(1, 2))
	fmt.Println(sumVariable(1, 2, 3))
	fmt.Println(sumVariable(1, 2, 3, 4))
	printNames("Ana")
	printNames("Ana", "Bob")
	printNames("Ana", "Bob", "Chandler")
	fmt.Println(getValues(2))
	// use only two values from function that returns 3
	double, _, quad := getValuesWithName(3)
	fmt.Println(double, quad)
}
