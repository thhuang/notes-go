package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := "Hello", "World"
	fmt.Println(a, b)
	fmt.Println(swap(a, b))
}
