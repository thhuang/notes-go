package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)

	b := []bool{true, true, false}
	fmt.Println(b)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
