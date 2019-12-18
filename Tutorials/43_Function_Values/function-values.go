package main

import (
	"fmt"
	"math"
)

func compute34(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypow := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypow(5, 12))
	fmt.Println(compute34(hypow))
	fmt.Println(compute34(math.Pow))
}
