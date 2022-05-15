package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

func parseArgs(args []string) (conf config, err error) {
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
	conf.imageGen, ok = imageGenerators[color]
	if !ok {
		return conf, fmt.Errorf(
			"wrong color: %s, should be one of: %s", color, colorAvail,
		)
	}

	return conf, nil
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
