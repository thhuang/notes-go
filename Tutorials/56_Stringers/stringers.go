package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	fmt.Println(person{"Arthur Dent", 42})
	fmt.Println(person{"Zaphod Beeblebrox", 9001})
}
