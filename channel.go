package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test(c chan int) {
	n := 0
	for {
		n++
		c <- n 
	}
}

func main() {
	c := make(chan int)
	go test(c)
	for i := range c {
		fmt.Println(i)
	}
}


