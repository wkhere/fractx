package main

type Fractal struct {
	w, h           int
	x0, y0, x1, y1 float64
	maxi           int
}

func (f *Fractal) Fill(image FractalImage) {
	dx := (f.x1 - f.x0) / float64(f.w)
	dy := (f.y1 - f.y0) / float64(f.h)
	for y := 0; y < f.h; y++ {
		cy := f.y0 + float64(y)*dy
		for x := 0; x < f.w; x++ {
			i := iter(f.x0+float64(x)*dx, cy, f.maxi)
			image.writePixel(x, y, i)
		}
	}
}
