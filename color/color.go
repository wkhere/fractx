package color

import (
	"encoding/hex"
	"fmt"
	"image/color"
)

var Names = map[string]string{
	"black":    "00",
	"white":    "ff",
	"yellow":   "ffff00",
	"orange":   "ff8000",
	"red":      "ff0000",
	"purple":   "800080",
	"blue":     "0000ff",
	"darkblue": "0000c0",
	"cyan":     "00ffff",
	"green":    "00ff00",
}

func DecodeHex(s string) (c color.Color, _ error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return c, err
	}
	switch len(b) {
	case 3:
		return color.RGBA{b[0], b[1], b[2], 0xff}, nil
	case 4:
		return color.RGBA{b[0], b[1], b[2], b[3]}, nil
	case 1:
		return color.Gray{b[0]}, nil
	default:
		return c, fmt.Errorf("expected 3-byte or 1-byte hex, got: %s", s)
	}
}

func Decode(s string) (color.Color, error) {
	x, ok := Names[s]
	if ok {
		s = x
	}
	return DecodeHex(s)
}

func MustDecode(s string) color.Color {
	c, err := Decode(s)
	if err != nil {
		panic(err)
	}
	return c
}

func EncodeHex(c color.Color) string {
	b := To8bit(c)
	return hex.EncodeToString(b[:])
}

func To8bit(c color.Color) (res [4]uint8) {
	r, g, b, a := c.RGBA()
	res[0] = uint8(r >> 8)
	res[1] = uint8(g >> 8)
	res[2] = uint8(b >> 8)
	res[3] = uint8(a >> 8)
	return
}
