package fractx

import (
	"image"
	"image/color"
	"sort"

	colorutil "github.com/wkhere/fractx/color"
)

type FractalImage interface {
	image.Image

	// writePixel writes pixel data for a given iteration.
	writePixel(x, y int, iter uint)
}

var ImageGenerators = map[string]func(*Fractal) FractalImage{
	"bw":   NewBWImage,
	"gray": NewGrayImage,
	"col1": curry(NewPalettedImage, Colorset1),
}

func ImageGeneratorNames() (ss []string) {
	ss = make([]string, len(ImageGenerators))
	i := 0
	for k := range ImageGenerators {
		ss[i] = k
		i++
	}
	sort.Strings(ss)
	return
}

type grayImage struct {
	*image.Gray
	maxi uint
	di   float64
}

func NewGrayImage(f *Fractal) FractalImage {
	return &grayImage{
		Gray: image.NewGray(image.Rect(0, 0, f.Size.W, f.Size.H)),
		maxi: f.MaxI,
		di:   256 / float64(f.MaxI),
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
		Gray: image.NewGray(image.Rect(0, 0, f.Size.W, f.Size.H)),
		maxi: f.MaxI,
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

func NewPalettedImage(cs Colorset, f *Fractal) FractalImage {
	var (
		colors    = make([]color.Color, len(cs)+1)
		itersteps = make([]uint, len(cs))
	)
	sort.Sort(cs)
	colors[0] = color.Black
	for j, v := range cs {
		itersteps[j] = f.MaxI * v.IterPercent / 100
		colors[j+1] = colorutil.MustDecode(v.Color)
	}

	return &palettedImage{
		Paletted: image.NewPaletted(
			image.Rect(0, 0, f.Size.W, f.Size.H),
			colors,
		),
		maxi:   f.MaxI,
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
