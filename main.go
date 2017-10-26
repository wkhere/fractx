package main

import (
	"flag"
	"image/png"
	"os"
)

func main() {
	img := newFractal(700, 400, -2.5, -1, 1, 1)

	filename := flag.String("f", "mandelbrot.png", "output file")
	flag.Parse()

	file, err := os.Create(*filename)
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
