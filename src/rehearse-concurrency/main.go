package main

import (
	"fmt"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {
	// channel
	g := 2
	fmt.Println(g)
	c := make(chan int)
	go doSomething(c)
	<-c
	// arrows
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}
