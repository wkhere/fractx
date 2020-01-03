package main

import (
	"flag"
	"fmt"
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
		die(fmt.Errorf("failed creating output file: %v", err))
	}

	defer func() {
		if err = w.Close(); err != nil {
			die(fmt.Errorf("failed closing output file: %v", err))
		}
	}()

	err = png.Encode(w, img)
	if err != nil {
		die(fmt.Errorf("failed writing to output file: %v", err))
	}

}

func fileFromName(s string) (io.WriteCloser, error) {
	if s == "-" {
		return os.Stdout, nil
	}
	return os.Create(s)
}

func die(err error) {
	fmt.Fprintln(os.Stderr, "fractx:", err)
	os.Exit(1)
}
