package main

import (
	"image"
)

type fractal struct {
	w, h           int
	x0, y0, x1, y1 float64
	dx, dy         float64
	fractalImage
}

func setup(w, h int, x0, y0, x1, y1 float64) fractal {
	return fractal{w, h, x0, y0, x1, y1,
		(x1 - x0) / float64(w),
		(y1 - y0) / float64(h),
		nil,
	}
}

func (f *fractal) fill() {
	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			i := iter(f.x0+float64(x)*f.dx, f.y0+float64(y)*f.dy)
			f.writePixel(x, y, i)
		}
	}
}

type fractalImage interface {
	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter int)
}

type FractalGray struct {
	fractal
	*image.Gray
	di float64
}

func (f *FractalGray) writePixel(x, y int, iter int) {
	pos := y*f.Stride + x
	f.Pix[pos] = -byte(float64(iter) * f.di)
}

func (f *FractalGray) Image() image.Image { return f.Gray }

func NewFractalGray(w, h int, x0, y0, x1, y1 float64) image.Image {
	f := &FractalGray{
		fractal: setup(w, h, x0, y0, x1, y1),
		Gray:    image.NewGray(image.Rect(0, 0, w, h)),
		di:      256 / float64(maxi),
	}
	f.fractalImage = f
	f.fill()
	return f
}

type FractalBW struct {
	fractal
	*image.Gray
}

func (f *FractalBW) writePixel(x, y int, iter int) {
	pos := y*f.Stride + x
	if iter < maxi {
		f.Pix[pos] = 255
	}
}

func NewFractalBW(w, h int, x0, y0, x1, y1 float64) image.Image {
	f := &FractalBW{
		fractal: setup(w, h, x0, y0, x1, y1),
		Gray:    image.NewGray(image.Rect(0, 0, w, h)),
	}
	f.fractalImage = f
	f.fill()
	return f
}

var maxi int = 200
