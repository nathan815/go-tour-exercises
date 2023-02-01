package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b, i := 0, 0, 0
	return func() int {
		res := 0
		if i == 1 {
			res = 1
		} else {
			res = a + b
		}
		a = b
		b = res
		i++
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

