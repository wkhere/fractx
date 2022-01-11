//go:build !amd64

package main

func iter(x0, y0 float64, maxi int) (i int) {
	x, y := x0, y0
	for i = 1; ; {
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
