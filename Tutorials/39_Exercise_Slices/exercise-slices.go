package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dx)
	for i := range img {
		row := make([]uint8, dy)
		for j := range row {
			row[j] = uint8(i^j)
		}
		img[i] = row
	}
	return img
}

func main() {
	pic.Show(Pic)
}

