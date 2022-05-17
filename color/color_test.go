package color

import (
	"image/color"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
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

	for i, tc := range tab {
		res, err := Decode(tc.s)
		if err != nil {
			t.Errorf("tc[%d] have unexpected error: %v", i, err)
			continue
		}

		var c1, c2 [4]uint32
		c1[0], c1[1], c1[2], c1[3] = res.RGBA()
		c2[0], c2[1], c2[2], c2[3] = tc.want.RGBA()
		if c1 != c2 {
			t.Errorf("tc[%d] mismatch\nhave %v\nwant %v", i, res, tc.want)
		}
	}
}
