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

func coloringNames() (ss []string) {
	ss = make([]string, len(imageGenerators))
	i := 0
	for k := range imageGenerators {
		ss[i] = k
		i++
	}
	return
}

func main() {
	var (
		maxi            int
		color, filename string
	)
	flag.IntVar(
		&maxi, "maxi",
		200,
		"max number of iterations",
	)
	flag.StringVar(
		&color, "color",
		"gray",
		"one of: "+strings.Join(coloringNames(), ", "),
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

	w, err := fileFromName(filename)
	if err != nil {
		die(fmt.Errorf("failed creating output file: %v", err))
	}

	defer func() {
		if err = w.Close(); err != nil {
			die(fmt.Errorf("failed closing output file: %v", err))
		}
	}()

	f := &Fractal{700, 400, -2.5, -1, 1, 1, maxi}

	img := imageGen(f)
	f.Fill(img)

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
