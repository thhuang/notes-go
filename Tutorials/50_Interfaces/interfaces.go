package main

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	x, y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a abser

	a = myFloat(-math.Pi)
	fmt.Println(a.abs())

	a = &vertex{3, 4}
	fmt.Println(a.abs())

}
