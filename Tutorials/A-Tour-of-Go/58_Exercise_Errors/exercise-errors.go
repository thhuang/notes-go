package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(-2)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(sqrt(-2))
	fmt.Println(math.Sqrt(-2))
}
