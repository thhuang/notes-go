package main

import "fmt"

func fibonacci() func() int {
	i, cache, fib := 0, 0, 0
	return func() int {
		if i == 1 {
			fib = 1
		} else {
			fib, cache = fib+cache, fib
		}
		i++
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
