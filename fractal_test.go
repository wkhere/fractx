package main

import (
	"testing"
)

var params = Params{700, 400, -2.5, -1, 1, 1, 200}

func BenchmarkFractalBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalBW(params)
	}
}

func BenchmarkFractalGray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFractalGray(params)
	}
}
