package main

import (
	"image"
	"image/color"
)

type fractal struct {
	w, h           int
	x0, y0, x1, y1 float64
	dx, dy         float64
	iterFunc       func(x0, y0 float64) int64
}

func newFractal(w, h int, x0, y0, x1, y1 float64) *fractal {
	return &fractal{
		w, h,
		x0, y0, x1, y1,
		(x1 - x0) / float64(w), (y1 - y0) / float64(h),
		iterSSE,
	}
}

func (_ fractal) ColorModel() color.Model { return color.RGBAModel }

func (img *fractal) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *fractal) At(x, y int) color.Color {
	i := img.iterFunc(img.x0+float64(x)*img.dx, img.y0+float64(y)*img.dy)
	if i >= maxi {
		return color.Black
	}
	return color.White
}

func iterSSE(x0, y0 float64) int64

var (
	maxi   int64   = 200
	pbound float64 = 4.0
	two    float64 = 2.0
)
