package main

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/spakin/netpbm"
)

func TestPPMContent(t *testing.T) {

	// In order to produce MD5 sum for the table below,
	// create fractx PNG and convert it via imagemagick to PPM:
	//   convert mandelbrot.png test.ppm
	// then md5sum PPM file.
	var tab = []struct {
		md5      string
		imageGen func(*Fractal) FractalImage
	}{
		{"25bb7c32464ff89ed9773ac479383f66", NewBWImage},
		{"0bf2783cd5d36fadaf9be441ba3afdbf", NewGrayImage},
	}

	ppmOpt := &netpbm.EncodeOptions{Format: netpbm.PPM}

	for i, tc := range tab {

		img := tc.imageGen(fractal)
		fractal.Fill(img)

		hash := md5.New()
		err := netpbm.Encode(hash, img, ppmOpt)
		if err != nil {
			t.Errorf("tc[%d] error during PPM Encode: %v", i, err)
			continue
		}

		md5 := hex.EncodeToString(hash.Sum(nil))

		if md5 != tc.md5 {
			t.Errorf("tc[%d] MD5 mismatch\nhave %s\nwant %s", i, md5, tc.md5)
		}

	}
}
