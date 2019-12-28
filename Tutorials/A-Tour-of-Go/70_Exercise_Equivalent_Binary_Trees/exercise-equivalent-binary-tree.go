package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Step(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Step(t.Left, ch)
	ch <- t.Value
	Step(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	Step(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	var t1, t2 *tree.Tree

	t1, t2 = tree.New(1), tree.New(1)
	fmt.Println(t1)
	fmt.Println(t2)
	if Same(t1, t2) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	t1, t2 = tree.New(1), tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)
	if Same(t1, t2) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
