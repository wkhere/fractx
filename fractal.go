package main

type Size struct {
	w, h int
}

type Rect struct {
	x0, y0, x1, y1 float64
}

type Fractal struct {
	size   Size
	bounds Rect
	maxi   uint
}

func (f *Fractal) Fill(image FractalImage) {
	dx := (f.bounds.x1 - f.bounds.x0) / float64(f.size.w)
	dy := (f.bounds.y1 - f.bounds.y0) / float64(f.size.h)
	for y := 0; y < f.size.h; y++ {
		cy := f.bounds.y0 + float64(y)*dy
		for x := 0; x < f.size.w; x++ {
			i := iter(f.bounds.x0+float64(x)*dx, cy, f.maxi)
			image.writePixel(x, y, i)
		}
	}
}
