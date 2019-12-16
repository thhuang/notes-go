package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 3
	fmt.Println(v)

	p := &v
	p.Y = 4
	fmt.Println(v)
}
