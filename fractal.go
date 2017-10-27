package main

import (
	"image"
	"sync"
)

type Fractal struct {
	w, h           int
	x0, y0, x1, y1 float64
	dx, dy         float64
	fractalImage
}

func setup(w, h int, x0, y0, x1, y1 float64) Fractal {
	return Fractal{w, h, x0, y0, x1, y1,
		(x1 - x0) / float64(w),
		(y1 - y0) / float64(h),
		nil,
	}
}

func (f *Fractal) fill() {
	var wg sync.WaitGroup
	wg.Add(4)

	chunk := func(py0, py1 int) {
		for py := py0; py < f.h; py++ {
			for px := 0; px < f.w; px++ {
				i := iter(f.x0+float64(px)*f.dx, f.y0+float64(py)*f.dy)
				f.writePixel(px, py, i)
			}
		}
		wg.Done()
	}
	go chunk(0, f.h/4)
	go chunk(f.h/4+1, f.h*2)
	go chunk(f.h/2+1, f.h*3/4)
	go chunk(f.h*3/4+1, f.h)

	wg.Wait()
}

type fractalImage interface {
	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter int)
}

type FractalGray struct {
	Fractal
	*image.Gray
	di float64
}

func (f *FractalGray) writePixel(x, y int, iter int) {
	pos := y*f.Stride + x
	f.Pix[pos] = -byte(float64(iter) * f.di)
}

func NewFractalGray(w, h int, x0, y0, x1, y1 float64) (r *FractalGray) {
	r = &FractalGray{
		Fractal: setup(w, h, x0, y0, x1, y1),
		Gray:    image.NewGray(image.Rect(0, 0, w, h)),
		di:      256 / float64(maxi),
	}
	r.fractalImage = r
	r.fill()
	return
}

type FractalBW struct {
	Fractal
	*image.Gray
}

func (f *FractalBW) writePixel(x, y int, iter int) {
	pos := y*f.Stride + x
	if iter < maxi {
		f.Pix[pos] = 255
	}
}

func NewFractalBW(w, h int, x0, y0, x1, y1 float64) (r *FractalBW) {
	r = &FractalBW{
		Fractal: setup(w, h, x0, y0, x1, y1),
		Gray:    image.NewGray(image.Rect(0, 0, w, h)),
	}
	r.fractalImage = r
	r.fill()
	return
}

var maxi int = 200
