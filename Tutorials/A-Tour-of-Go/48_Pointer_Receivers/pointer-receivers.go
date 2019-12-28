package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *vertex) scale(f float64) {
	v.x *= f
	v.y *= f
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v, v.abs())

	v.scale(10)
	fmt.Println(v, v.abs())

	p := &v
	p.scale(10)
	fmt.Println(p, p.abs())
}
