//go:build !amd64

package fractx

func iter(x0, y0 float64, maxi uint) (i uint) {
	x, y := x0, y0
	for i = 0; ; {
		xx, yy := x*x, y*y
		if xx+yy >= 4 {
			break
		}
		x, y = xx-yy+x0, 2*x*y+y0
		if i++; i >= maxi {
			break
		}
	}
	return
}
