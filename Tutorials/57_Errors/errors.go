package main

import (
	"fmt"
	"time"
)

type myError struct {
	time    time.Time
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v, %v", e.time, e.message)
}

func run() error {
	return &myError{
		time:    time.Now(),
		message: "Something went wrong!",
	}
}

func main() {
	fmt.Println(run())
}
