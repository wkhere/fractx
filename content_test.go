package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/spakin/netpbm"
)

func TestPPMContent(t *testing.T) {

	// In order to produce MD5 sum for the table below,
	// create fractx PNG and convert it via imagemagick to PPM:
	//   convert mandelbrot.png test.ppm
	// then md5sum PPM file.
	type tcdata struct {
		md5   string
		color string
		maxi  int
	}
	var tab = []tcdata{
		{"25bb7c32464ff89ed9773ac479383f66", "bw", 200},
		{"599a7caed7d1bce58a3d37e0127e4278", "bw", 255},
		{"59567e6b2808da4d2fb8be3b962e9d54", "bw", 256},
		{"c07a91142c7b1a81db984b0f8ba927b2", "bw", 257},
		{"fb4559ac9da7a646b1c0c2edf22db01c", "gray", 255},
		{"0bf2783cd5d36fadaf9be441ba3afdbf", "gray", 200},
		{"fb4559ac9da7a646b1c0c2edf22db01c", "gray", 256},
		{"53ae6f53698470da682eb142e0f713cd", "gray", 257},
	}

	ppmOpt := &netpbm.EncodeOptions{Format: netpbm.PPM}

	nameGen := func(tc *tcdata) string {
		return fmt.Sprintf("%s-%d", tc.color, tc.maxi)
	}

	for _, tc := range tab {
		tc := tc // pfff
		t.Run(nameGen(&tc), func(t *testing.T) {
			t.Parallel()

			fractal := &Fractal{700, 400, -2.5, -1, 1, 1, tc.maxi}

			img := imageGenerators[tc.color](fractal)
			fractal.Fill(img)

			hash := md5.New()
			err := netpbm.Encode(hash, img, ppmOpt)
			if err != nil {
				t.Errorf("error during PPM Encode: %v", err)
				return
			}

			md5 := hex.EncodeToString(hash.Sum(nil))

			if md5 != tc.md5 {
				t.Errorf("MD5 mismatch\nhave %s\nwant %s", md5, tc.md5)
			}
		})
	}
}
