package main

import (
	"image/png"
	"os"
)

func main() {

	img := newFractal(700, 400, -2.5, -1, 1, 1)

	file, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	png.Encode(file, img)

}
