package main

import (
	"image/png"
	"io/ioutil"
	"testing"
)

var img = newFractal(700, 400, -2.5, -1, 1, 1)

func BenchmarkImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		png.Encode(ioutil.Discard, img)
	}
}
