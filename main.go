package main

import (
	"flag"
	"image"
	"image/png"
	"os"
)

func main() {
	var img image.Image

	bw := flag.Bool("bw", false, "black&white")
	filename := flag.String("f", "mandelbrot.png", "output file")
	flag.Parse()

	switch {
	case *bw:
		img = NewFractalBW(700, 400, -2.5, -1, 1, 1)
	default:
		img = NewFractalGray(700, 400, -2.5, -1, 1, 1)
	}

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
