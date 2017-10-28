package main

import (
	"image"
)

type Params struct {
	w, h           int
	x0, y0, x1, y1 float64
}

type fractal struct {
	Params
	dx, dy float64
	fractalImage
}

func setup(p *Params) (f fractal) {
	f.Params = *p
	f.dx = (p.x1 - p.x0) / float64(p.w)
	f.dy = (p.y1 - p.y0) / float64(p.h)
	return
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

func NewFractalGray(params Params) image.Image {
	f := &FractalGray{
		fractal: setup(&params),
		Gray:    image.NewGray(image.Rect(0, 0, params.w, params.h)),
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

func NewFractalBW(params Params) image.Image {
	f := &FractalBW{
		fractal: setup(&params),
		Gray:    image.NewGray(image.Rect(0, 0, params.w, params.h)),
	}
	f.fractalImage = f
	f.fill()
	return f
}

var maxi int = 200
