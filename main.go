package main

import (
	"fmt"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const prog = "fractx"

type config struct {
	maxi     int
	imageGen func(*Fractal) FractalImage
	filename string
}

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

func parseArgs(args []string) (conf config) {
	var (
		colorAvail = strings.Join(coloringNames(), ", ")
		color      string
		help       bool
	)

	flag := pflag.NewFlagSet("flags", pflag.ContinueOnError)
	flag.SortFlags = false

	flag.IntVarP(&conf.maxi, "maxi", "i", 200,
		"max number of iterations")
	flag.StringVarP(&color, "color", "c", "gray",
		"one of: "+colorAvail)
	flag.StringVarP(&conf.filename, "output", "o", "mandelbrot.png",
		"output file or '-' for stdout")

	flag.BoolVarP(&help, "help", "h", false,
		"show this help and exit")

	flag.Usage = func() {
		fmt.Fprintln(flag.Output(), "Generate Mandelbrot fractal.")
		fmt.Fprintln(flag.Output(), "Usage:", prog, "[FLAGS]")
		flag.PrintDefaults()
	}

	err := flag.Parse(args)
	if err != nil {
		die(2, err)
	}
	if help {
		flag.SetOutput(os.Stdout)
		flag.Usage()
		die(0)
	}

	var ok bool
	conf.imageGen, ok = imageGenerators[color]
	if !ok {
		die(2, "wrong color:", color+", should be one of:", colorAvail)
	}

	return conf
}

func main() {
	conf := parseArgs(os.Args[1:])

	w, err := fileFromName(conf.filename)
	if err != nil {
		die(1, "failed creating output file:", err)
	}

	defer func() {
		if err = w.Close(); err != nil {
			die(1, "failed closing output file:", err)
		}
	}()

	f := &Fractal{700, 400, -2.5, -1, 1, 1, conf.maxi}

	img := conf.imageGen(f)
	f.Fill(img)

	err = png.Encode(w, img)
	if err != nil {
		die(1, "failed writing to output file:", err)
	}

}

func fileFromName(s string) (io.WriteCloser, error) {
	if s == "-" {
		return os.Stdout, nil
	}
	return os.Create(s)
}

func die(exitcode int, msgs ...interface{}) {
	if len(msgs) > 0 {
		fmt.Fprint(os.Stderr, prog, ": ")
		fmt.Fprintln(os.Stderr, msgs...)
	}
	os.Exit(exitcode)
}
