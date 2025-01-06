package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wkhere/fractx"
)

const prog = "fractx"

type config struct {
	size      fractx.Size
	bounds    fractx.Rect
	maxi      uint
	newImage  fractx.ImageBuilder
	filename  string
	overwrite bool
	web       bool
	webAddr   string

	help func(io.Writer)
}

func run(c *config) error {
	f := &fractx.Fractal{Size: c.size, Bounds: c.bounds, MaxI: c.maxi}

	if c.web {
		return (&server{
			fractal:  f,
			origSize: c.size,
			newImage: c.newImage,
			addr:     c.webAddr,
		}).serve()
	} else {
		return (&dumper{
			fractal:   f,
			newImage:  c.newImage,
			filename:  c.filename,
			overwrite: c.overwrite,
		}).dump()
	}
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

func die(exitcode int, msgs ...interface{}) {
	if len(msgs) > 0 {
		fmt.Fprint(os.Stderr, prog, ": ")
		fmt.Fprintln(os.Stderr, msgs...)
	}
	os.Exit(exitcode)
}
