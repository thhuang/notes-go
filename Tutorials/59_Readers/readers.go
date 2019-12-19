package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf(
			"n: %v\nerr: %v\nb: %v\nb[:n]: %q\n\n",
			n, err, b, b[:n],
		)
		if err == io.EOF {
			break
		}
	}
}
