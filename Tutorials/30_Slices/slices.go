package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	a := primes[1:4]
	b := primes[2:6]
	fmt.Println(a, b)

	b[0] = 17
	fmt.Println(a, b)
}
