package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type pixel struct {
	r, g, b, a uint8
}

type myImage struct {
	pix  [][]pixel
	rect image.Rectangle
}

func (p *myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *myImage) Bounds() image.Rectangle {
	return p.rect
}

func (p *myImage) At(x, y int) color.Color {
	return color.RGBA{
		p.pix[x][y].r,
		p.pix[x][y].g,
		p.pix[x][y].b,
		p.pix[x][y].a,
	}
}

func newMyImage(r image.Rectangle) *myImage {
	x, y := r.Dx(), r.Dy()
	pix := make([][]pixel, x)
	for i := range pix {
		pix[i] = make([]pixel, y)
	}
	return &myImage{pix, r}
}

func myPaint(p *myImage) {
	for i := 0; i < p.Bounds().Dx(); i++ {
		for j := 0; j < p.Bounds().Dy(); j++ {
			c := uint8(i ^ j)
			p.pix[i][j] = pixel{c, c, c, 255}
		}
	}
}

func main() {
	m := newMyImage(image.Rect(0, 0, 256, 256))
	myPaint(m)
	pic.ShowImage(m)
}
