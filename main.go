package main

import (
	"flag"
	"image/png"
	"io"
	"os"
	"strings"
)

var imageGenerators = map[string]func(*Fractal) FractalImage{
	"bw":   NewBWImage,
	"gray": NewGrayImage,
}

var coloringNames []string

func init() {
	coloringNames = make([]string, len(imageGenerators))
	i := 0
	for k := range imageGenerators {
		coloringNames[i] = k
		i++
	}
}

func main() {
	var (
		color, filename string
	)
	flag.StringVar(
		&color, "color",
		"gray",
		"one of: "+strings.Join(coloringNames, ", "),
	)
	flag.StringVar(
		&filename, "o",
		"mandelbrot.png",
		"output file or '-' for stdout",
	)
	flag.Parse()

	imageGen, ok := imageGenerators[color]
	if !ok {
		flag.Usage()
		os.Exit(2)
	}

	f := &Fractal{700, 400, -2.5, -1, 1, 1, 200}

	img := imageGen(f)
	f.Fill(img)

	w, err := fileFromName(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}()

	png.Encode(w, img)
}

func fileFromName(s string) (io.WriteCloser, error) {
	if s == "-" {
		return os.Stdout, nil
	}
	return os.Create(s)
}
