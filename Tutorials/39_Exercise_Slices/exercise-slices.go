package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var img [][]uint8
	for i := 0; i < dx; i++ {
		var row []uint8
		for j := 0; j < dy; j++ {
			row = append(row, uint8(i^j))
		}
		img = append(img, row)
	}
	return img
}

func main() {
	pic.Show(Pic)
}

