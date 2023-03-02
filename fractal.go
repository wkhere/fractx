package fractx

type Size struct {
	W, H int
}

type Rect struct {
	X0, Y0, X1, Y1 float64
}

type Fractal struct {
	Size   Size
	Bounds Rect
	MaxI   uint
}

func (f *Fractal) Fill(image Image) {
	dx := (f.Bounds.X1 - f.Bounds.X0) / float64(f.Size.W)
	dy := (f.Bounds.Y1 - f.Bounds.Y0) / float64(f.Size.H)
	for y := 0; y < f.Size.H; y++ {
		cy := f.Bounds.Y0 + float64(y)*dy
		for x := 0; x < f.Size.W; x++ {
			i := iter(f.Bounds.X0+float64(x)*dx, cy, f.MaxI)
			image.writePixel(x, y, i)
		}
	}
}
