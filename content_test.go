package fractx

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
		maxi  uint
	}
	var tab = []tcdata{
		{"5f37292662ca4b5c5b76bf57dd994aab", "bw", 200},
		{"59567e6b2808da4d2fb8be3b962e9d54", "bw", 255},
		{"c07a91142c7b1a81db984b0f8ba927b2", "bw", 256},
		{"0e80a022e27d45758defd3626c31c843", "bw", 257},
		{"6a7166e93bb7b91c66bb418087876329", "gray", 200},
		{"53ae6f53698470da682eb142e0f713cd", "gray", 255},
		{"53ae6f53698470da682eb142e0f713cd", "gray", 256},
		{"6f3fc32225cf5a284d2d11fdc73f3deb", "gray", 257},
		{"2b44d789990562077ecfd196ecec1474", "col1", 100},
	}

	ppmOpt := &netpbm.EncodeOptions{Format: netpbm.PPM}

	nameGen := func(tc *tcdata) string {
		return fmt.Sprintf("%s-%d", tc.color, tc.maxi)
	}

	for _, tc := range tab {
		tc := tc // pfff
		t.Run(nameGen(&tc), func(t *testing.T) {
			t.Parallel()

			fractal := &Fractal{DefaultSize, DefaultBounds, tc.maxi}

			img := ImageGenerators[tc.color](fractal)
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
