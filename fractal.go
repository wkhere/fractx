package main

import (
	"image"
)

type Fractal struct {
	w, h           int
	x0, y0, x1, y1 float64
	maxi           int
}

func (f *Fractal) Fill(image FractalImage) {
	dx := (f.x1 - f.x0) / float64(f.w)
	dy := (f.y1 - f.y0) / float64(f.h)
	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			i := iter(f.x0+float64(x)*dx, f.y0+float64(y)*dy, f.maxi)
			image.writePixel(x, y, i)
		}
	}
}

type FractalImage interface {
	image.Image
	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter int)
}

type grayImage struct {
	*image.Gray
	di float64
}

func (img *grayImage) writePixel(x, y int, iter int) {
	pos := y*img.Stride + x
	img.Pix[pos] = -byte(float64(iter) * img.di)
}

func NewGrayImage(f *Fractal) FractalImage {
	return &grayImage{
		Gray: image.NewGray(image.Rect(0, 0, f.w, f.h)),
		di:   256 / float64(f.maxi),
	}
}

type bwImage struct {
	*image.Gray
	maxi int
}

func (img *bwImage) writePixel(x, y int, iter int) {
	pos := y*img.Stride + x
	if iter < img.maxi {
		img.Pix[pos] = 255
	}
}

func NewBWImage(f *Fractal) FractalImage {
	return &bwImage{
		Gray: image.NewGray(image.Rect(0, 0, f.w, f.h)),
		maxi: f.maxi,
	}
}
