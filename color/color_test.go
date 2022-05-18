package color

import (
	"image/color"
	"testing"
)

var tab = []struct {
	s    string
	want color.Color
}{
	{"black", color.Black},
	{"white", color.White},
	{"000000", color.Black},
	{"00", color.Black},
	{"000000ff", color.Black},
	{"ffffffff", color.White},
}

func TestDecodeBasic(t *testing.T) {
	for i, tc := range tab {
		res, err := Decode(tc.s)
		if err != nil {
			t.Errorf("tc[%d] have unexpected error: %v", i, err)
			continue
		}

		c1 := To8bit(tc.want)
		c2 := To8bit(res)
		if c1 != c2 {
			t.Errorf("tc[%d] mismatch\nhave %v\nwant %v", i, c2, c1)
		}
	}
}

func FuzzEncodeDecodeHex(f *testing.F) {
	f.Add(byte(0), byte(0), byte(0), byte(0xff))
	f.Add(byte(0x80), byte(0xfc), byte(0xe2), byte(0x80))
	f.Add(byte(0xd0), byte(0x10), byte(0x20), byte(0xff))
	f.Add(byte(0xff), byte(0xff), byte(0xff), byte(0xff))
	f.Add(byte(0xc0), byte(0xc0), byte(0xc0), byte(0))

	f.Fuzz(func(t *testing.T, r, g, b, a byte) {
		c := color.RGBA{r, g, b, a}

		s := EncodeHex(c)

		res, err := DecodeHex(s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		c1 := To8bit(c)
		c2 := To8bit(res)
		if c1 != c2 {
			t.Errorf("mismatch\nhave %v\nwant %v", c2, c1)
		}
	})
}
