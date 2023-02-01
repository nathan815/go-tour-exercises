package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z, last_z := 1.0, 0.0
	iterations := 0
	for math.Abs(z-last_z) > 0.0000000001 {
		last_z = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		iterations++
	}
	return z, iterations
}

func main() {
	x := 165335436.0
	my_sqrt, iterations := Sqrt(x)
	sqrt := math.Sqrt(x)
	fmt.Println("Mine=", my_sqrt, " Stdlib=", sqrt, "Diff=", math.Abs(sqrt-my_sqrt))
	fmt.Println(iterations, "Iterations")
	if iterations > 10 {
		fmt.Println("More than 10 iterations!")
	} else {
		fmt.Println("Less than 10 :)")
	}
}

