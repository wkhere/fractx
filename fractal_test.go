package main

import (
	"testing"
)

func BenchmarkFractalBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalBW(700, 400, -2.5, -1, 1, 1)
	}
}

func BenchmarkFractalGray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalGray(700, 400, -2.5, -1, 1, 1)
	}
}
