package main

import (
	"testing"
)

func BenchmarkImageBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalBW(700, 400, -2.5, -1, 1, 1)
	}
}

func BenchmarkImageGS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalGray(700, 400, -2.5, -1, 1, 1)
	}
}
