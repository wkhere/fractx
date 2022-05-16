package main

import (
	"encoding/hex"
	"fmt"
	"image/color"
)

var colornames = map[string]string{
	"black":    "000000",
	"white":    "ffffff",
	"yellow":   "ffff00",
	"orange":   "ff8000",
	"red":      "ff0000",
	"purple":   "800080",
	"blue":     "0000ff",
	"darkblue": "0000c0",
	"cyan":     "00ffff",
	"green":    "00ff00",
}

func DecodeHexColor(s string) (c color.Color, _ error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return c, err
	}
	switch {
	case len(b) == 3:
		return color.RGBA{b[0], b[1], b[2], 0xff}, nil
	case len(b) == 1:
		return color.Gray{b[0]}, nil
	default:
		return c, fmt.Errorf("expected 3-byte or 1-byte hex, got: %s", s)
	}
}

func DecodeColor(s string) (color.Color, error) {
	x, ok := colornames[s]
	if ok {
		s = x
	}
	return DecodeHexColor(s)
}

func MustDecodeColor(s string) color.Color {
	c, err := DecodeColor(s)
	if err != nil {
		panic(err)
	}
	return c
}
