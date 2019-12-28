package main

import (
	"fmt"
	"time"
)

func bomb(tick, boom <-chan time.Time) {
	for {
		select {
		case <-tick:
			fmt.Print("tick")
		case <-boom:
			fmt.Print("\nBOOM!\n")
			return
		default:
			time.Sleep(10 * time.Millisecond)
			fmt.Print(".")
		}
	}
}

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(time.Second)

	bomb(tick, boom)
}
