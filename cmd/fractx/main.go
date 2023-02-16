package main

import (
	"fmt"
	"image/png"
	"io"
	"os"

	fx "github.com/wkhere/fractx"
)

const prog = "fractx"

type config struct {
	size      fx.Size
	bounds    fx.Rect
	maxi      uint
	imageGen  func(*fx.Fractal) fx.FractalImage
	filename  string
	overwrite bool

	help func(io.Writer)
}

func run(c *config) (err error) {
	w, err := fileFromName(c.filename, c.overwrite)
	if err != nil {
		return fmt.Errorf("failed creating output file: %w", err)
	}
	defer func() {
		if cerr := w.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed closing output file: %w", err)
		}
	}()

	f := &fx.Fractal{c.size, c.bounds, c.maxi}

	img := c.imageGen(f)
	f.Fill(img)

	err = png.Encode(w, img)
	if err != nil {
		return fmt.Errorf("failed writing to output file: %w", err)
	}
	return nil
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

	err = run(&conf)
	if err != nil {
		die(1, err)
	}
}

func fileFromName(s string, overwrite bool) (*os.File, error) {
	if s == "-" {
		return os.Stdout, nil
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	if !overwrite {
		flag |= os.O_EXCL
	}
	return os.OpenFile(s, flag, 0644)
}

func die(exitcode int, msgs ...interface{}) {
	if len(msgs) > 0 {
		fmt.Fprint(os.Stderr, prog, ": ")
		fmt.Fprintln(os.Stderr, msgs...)
	}
	os.Exit(exitcode)
}
