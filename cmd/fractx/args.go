package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	fx "github.com/wkhere/fractx"
)

func parseArgs(args []string) (conf config, err error) {
	var (
		colorAvail = strings.Join(fx.ImageGeneratorNames(), ", ")
		color      string
		help       bool
	)
	flag := pflag.NewFlagSet("flags", pflag.ContinueOnError)
	flag.SortFlags = false

	conf.size = fx.DefaultSize
	conf.bounds = fx.DefaultBounds
	flag.VarP(&conf.size, "size", "s", "pixel size")
	flag.VarP(&conf.bounds, "bounds", "b", "float-plane bounds")

	flag.UintVarP(&conf.maxi, "maxi", "i", fx.DefaultMaxI,
		"max number of iterations")
	flag.StringVarP(&color, "color", "c", "col1",
		"one of: "+colorAvail)
	flag.StringVarP(&conf.filename, "output", "o", "mandelbrot.png",
		"output file or '-' for stdout")
	flag.BoolVarP(&conf.overwrite, "overwrite", "O", true,
		"overwrite output file")

	flag.BoolVarP(&help, "help", "h", false,
		"show this help and exit")

	flag.Usage = func() {
		fmt.Fprintln(flag.Output(), "Generate Mandelbrot fractal.")
		fmt.Fprintln(flag.Output(), "Usage:", prog, "[FLAGS]")
		flag.PrintDefaults()
	}

	err = flag.Parse(args)
	if err != nil {
		return conf, err
	}
	if help {
		conf.help = func(w io.Writer) {
			flag.SetOutput(w)
			flag.Usage()
		}
		return conf, nil
	}

	var ok bool
	conf.imageGen, ok = fx.ImageGenerators[color]
	if !ok {
		return conf, fmt.Errorf(
			"wrong color: %s, should be one of: %s", color, colorAvail,
		)
	}

	return conf, nil
}
