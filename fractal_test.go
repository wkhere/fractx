package main

import (
	"testing"
)

var fractal = &Fractal{DefaultSize, DefaultBounds, DefaultMaxI}

func BenchmarkFractalBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := NewBWImage(fractal)
		fractal.Fill(img)
	}
}

func BenchmarkFractalGray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := NewGrayImage(fractal)
		fractal.Fill(img)
	}
}

func BenchmarkFractalCol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := NewPalettedImage(colorset1, fractal)
		fractal.Fill(img)
	}
}
