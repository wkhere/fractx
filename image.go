package main

import (
	"image"
	"image/color"
	"sort"
)

type FractalImage interface {
	image.Image

	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter uint)
}

type grayImage struct {
	*image.Gray
	maxi uint
	di   float64
}

func NewGrayImage(f *Fractal) FractalImage {
	return &grayImage{
		Gray: image.NewGray(image.Rect(0, 0, f.size.w, f.size.h)),
		maxi: f.maxi,
		di:   256 / float64(f.maxi),
	}
}

func (img *grayImage) writePixel(x, y int, iter uint) {
	if iter < img.maxi {
		pos := y*img.Stride + x
		img.Pix[pos] = 255 - byte(float64(iter)*img.di)
	}
}

type bwImage struct {
	*image.Gray
	maxi uint
}

func NewBWImage(f *Fractal) FractalImage {
	return &bwImage{
		Gray: image.NewGray(image.Rect(0, 0, f.size.w, f.size.h)),
		maxi: f.maxi,
	}
}

func (img *bwImage) writePixel(x, y int, iter uint) {
	if iter < img.maxi {
		pos := y*img.Stride + x
		img.Pix[pos] = 255
	}
}

type palettedImage struct {
	*image.Paletted
	maxi   uint
	isteps []uint
}

func NewPalettedImage(cs colorset, f *Fractal) FractalImage {
	var (
		colors    = make([]color.Color, len(cs)+1)
		itersteps = make([]uint, len(cs))
	)
	sort.Sort(cs)
	colors[0] = color.Black
	for j, v := range cs {
		itersteps[j] = f.maxi * v.iterPercent / 100
		colors[j+1] = MustDecodeColor(v.color)
	}

	return &palettedImage{
		Paletted: image.NewPaletted(
			image.Rect(0, 0, f.size.w, f.size.h),
			colors,
		),
		maxi:   f.maxi,
		isteps: itersteps,
	}
}

func (img *palettedImage) writePixel(x, y int, iter uint) {
	if iter >= img.maxi {
		return
	}
	var j uint8
	for ; j < uint8(len(img.isteps)); j++ {
		if iter >= img.isteps[j] {
			break
		}
	}
	pos := y*img.Stride + x
	img.Pix[pos] = j + 1
}
