package main

import (
	"image/png"
	"io/ioutil"
	"testing"
)

var imgBW = NewFractalBW(700, 400, -2.5, -1, 1, 1)
var imgGS = NewFractalGray(700, 400, -2.5, -1, 1, 1)

func BenchmarkImageBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		png.Encode(ioutil.Discard, imgBW)
	}
}

func BenchmarkImageGS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		png.Encode(ioutil.Discard, imgGS)
	}
}
