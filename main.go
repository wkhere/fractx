package main

import (
	"flag"
	"image"
	"image/png"
	"os"
	"strings"
)

var fractals = map[string]func(Params) image.Image{
	"bw":   NewFractalBW,
	"gray": NewFractalGray,
}

var coloringNames []string

func init() {
	coloringNames = make([]string, len(fractals))
	i := 0
	for k := range fractals {
		coloringNames[i] = k
		i++
	}
}

func main() {
	color := flag.String("color", "gray", "one of: "+
		strings.Join(coloringNames, ", "))
	filename := flag.String("o", "mandelbrot.png", "output file")
	flag.Parse()

	f, ok := fractals[*color]
	if !ok {
		flag.Usage()
		os.Exit(2)
	}
	img := f(Params{700, 400, -2.5, -1, 1, 1})

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
