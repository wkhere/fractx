package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/wkhere/fractx"
)

type dumper struct {
	fractal   *fractx.Fractal
	newImage  fractx.ImageBuilder
	filename  string
	overwrite bool
}

func (d *dumper) dump() (err error) {
	w, err := fileFromName(d.filename, d.overwrite)
	if err != nil {
		return fmt.Errorf("failed creating output file: %w", err)
	}
	defer func() {
		if cerr := w.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed closing output file: %w", err)
		}
	}()

	img := d.newImage(d.fractal)
	d.fractal.Fill(img)

	err = png.Encode(w, img)
	if err != nil {
		return fmt.Errorf("failed writing to output file: %w", err)
	}
	return nil
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
