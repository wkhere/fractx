package main

import (
	"fmt"
	"io"
	"strconv"
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

	conf.size = DefaultSize
	conf.bounds = DefaultBounds
	flag.VarP(&conf.size, "size", "s", "pixel size")
	flag.VarP(&conf.bounds, "bounds", "b", "float-plane bounds")

	flag.IntVarP(&conf.maxi, "maxi", "i", DefaultMaxI,
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

func (s *Size) Set(input string) (err error) {
	cc := strings.Split(input, ",")
	if len(cc) != 2 {
		return fmt.Errorf("expected width,height")
	}
	s.w, err = strconv.Atoi(strings.TrimSpace(cc[0]))
	if err != nil {
		return err
	}
	if s.w <= 0 {
		return fmt.Errorf(`parsed "%d": width <= 0`, s.w)
	}
	s.h, err = strconv.Atoi(strings.TrimSpace(cc[1]))
	if err != nil {
		return err
	}
	if s.h <= 0 {
		return fmt.Errorf(`parsed "%d": height <= 0`, s.h)
	}
	return nil
}

func (s *Size) String() string {
	return fmt.Sprintf("%d,%d", s.w, s.h)
}

func (*Size) Type() string {
	return "width,height"
}

func (r *Rect) Set(input string) (err error) {
	cc := strings.Split(input, ",")
	if len(cc) != 4 {
		return fmt.Errorf("expected x0,y0,x1,y1")
	}
	r.x0, err = strconv.ParseFloat(strings.TrimSpace(cc[0]), 64)
	if err != nil {
		return err
	}
	r.y0, err = strconv.ParseFloat(strings.TrimSpace(cc[1]), 64)
	if err != nil {
		return err
	}
	r.x1, err = strconv.ParseFloat(strings.TrimSpace(cc[2]), 64)
	if err != nil {
		return err
	}
	r.y1, err = strconv.ParseFloat(strings.TrimSpace(cc[3]), 64)
	if err != nil {
		return err
	}
	return nil
}

func (r *Rect) String() string {
	return fmt.Sprintf("%g,%g,%g,%g", r.x0, r.y0, r.x1, r.y1)
}

func (*Rect) Type() string {
	return "x0,y0,x1,y1"
}
