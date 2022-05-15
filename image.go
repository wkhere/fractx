package main

import "image"

type FractalImage interface {
	image.Image

	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter int)
}

type grayImage struct {
	*image.Gray
	maxi int
	di   float64
}

func NewGrayImage(f *Fractal) FractalImage {
	return &grayImage{
		Gray: image.NewGray(image.Rect(0, 0, f.size.w, f.size.h)),
		maxi: f.maxi,
		di:   256 / float64(f.maxi),
	}
}

func (img *grayImage) writePixel(x, y int, iter int) {
	pos := y*img.Stride + x
	if iter < img.maxi {
		img.Pix[pos] = 255 - byte(float64(iter)*img.di)
	}
}

type bwImage struct {
	*image.Gray
	maxi int
}

func NewBWImage(f *Fractal) FractalImage {
	return &bwImage{
		Gray: image.NewGray(image.Rect(0, 0, f.size.w, f.size.h)),
		maxi: f.maxi,
	}
}

func (img *bwImage) writePixel(x, y int, iter int) {
	pos := y*img.Stride + x
	if iter < img.maxi {
		img.Pix[pos] = 255
	}
}
