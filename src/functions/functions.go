package main

import (
	"fmt"
	"time"
)

func main() {
	// anonymous function
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)
	// anonymous with go routine
	c := make(chan int)
	go func() {
		fmt.Println("Start function")
		time.Sleep(5 * time.Second)
		fmt.Println("End function")
		c <- 1
	}()
	<-c

}
