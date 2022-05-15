package main

import (
	"testing"
)

var fractal = &Fractal{Size{700, 400}, Rect{-2.5, -1, 1, 1}, 200}

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
