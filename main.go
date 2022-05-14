package main

import (
	"fmt"
	"image/png"
	"io"
	"os"
)

const prog = "fractx"

type config struct {
	maxi     int
	imageGen func(*Fractal) FractalImage
	filename string

	help func(io.Writer)
}

var imageGenerators = map[string]func(*Fractal) FractalImage{
	"bw":   NewBWImage,
	"gray": NewGrayImage,
}

func main() {
	conf, err := parseArgs(os.Args[1:])
	if err != nil {
		die(2, err)
	}
	if conf.help != nil {
		conf.help(os.Stdout)
		die(0)
	}

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
