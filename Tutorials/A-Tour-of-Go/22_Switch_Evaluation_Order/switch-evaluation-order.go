package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
