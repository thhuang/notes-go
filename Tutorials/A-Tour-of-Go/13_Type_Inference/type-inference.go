package main

import "fmt"

func main() {
	i := 12
	f := 3.14
	g := 1.2 + 3.4i
	fmt.Printf("i is of type %T, %v\n", i, i)
	fmt.Printf("f is of type %T, %v\n", f, f)
	fmt.Printf("g is of type %T, %v\n", g, g)
}
