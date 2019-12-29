package main

import "fmt"

var pc [256]byte

// pc[i] is the population count of i.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum int
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func main() {
	fmt.Println(PopCount(257))
}
