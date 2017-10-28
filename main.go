package main

import (
	"flag"
	"image"
	"image/png"
	"os"
)

type fractalF func(w, h int, x0, y0, x1, y1 float64) image.Image

var fractals = map[string]fractalF{
	"bw":   NewFractalBW,
	"gray": NewFractalGray,
}

func main() {
	color := flag.String("color", "gray", "one of: bw, gray")
	filename := flag.String("o", "mandelbrot.png", "output file")
	flag.Parse()

	f, ok := fractals[*color]
	if !ok {
		flag.Usage()
		os.Exit(2)
	}
	img := f(700, 400, -2.5, -1, 1, 1)

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
