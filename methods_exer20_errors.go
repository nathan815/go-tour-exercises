package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot take sqrt of negative %f", float64(e))
}

func Sqrt(x float64) (float64, int, error) {
	if (x < 0) {
		return 0, 0, ErrNegativeSqrt(x)
	}
	
	z, last_z := 1.0, 0.0
	iterations := 0
	for math.Abs(z-last_z) > 0.0000000001 {
		last_z = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		iterations++
	}
	return z, iterations, nil
}

func main() {
	x := 165335436.0
	my_sqrt, iterations, err := Sqrt(x)
	if err != nil {
		panic(err)
	}
	sqrt := math.Sqrt(x)
	fmt.Println("Mine=", my_sqrt, " Stdlib=", sqrt, "Diff=", math.Abs(sqrt-my_sqrt))
	fmt.Println(iterations, "Iterations")
	
	_, _, err = Sqrt(-5)
	if err != nil {
		fmt.Println(err)
	}
}

