package main

import "fmt"

const (
	big   = 1 << 200
	small = big >> 199
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}
