package main

import "fmt"

type vertex struct {
	lat, lon float64
}

func main() {
	m := make(map[string]vertex)
	m["Nokia Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Nokia Bell Labs"])
}
