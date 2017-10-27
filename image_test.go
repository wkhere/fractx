package main

import (
	"image/png"
	"io/ioutil"
	"testing"
)

func BenchmarkImageBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := NewFractalBW(700, 400, -2.5, -1, 1, 1)
		png.Encode(ioutil.Discard, img)
	}
}

func BenchmarkImageGS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := NewFractalGray(700, 400, -2.5, -1, 1, 1)
		png.Encode(ioutil.Discard, img)
	}
}
