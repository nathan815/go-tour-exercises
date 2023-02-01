package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
	fn     func(x, y int) int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x int, y int) color.Color {
	v := img.fn(x, y)
	return color.RGBA{uint8(v), uint8(v), uint8(v), 255}
}

func main() {
	m := Image{
		width:  500,
		height: 300,
		fn: func(x, y int) int {
			return (x*x + y*y) * 3
		},
	}
	pic.ShowImage(m)
}

