package main

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"
)

var img1, img2 *fractal

func init() {
	img1 = newFractal(700, 400, -2.5, -1, 1, 1)
	img2 = new(fractal)
	*img2 = *img1
	img1.iterFunc = iterSSE
	img2.iterFunc = iterGo
}

func TestImgEq(t *testing.T) {
	var b1, b2 bytes.Buffer
	png.Encode(&b1, img1)
	png.Encode(&b2, img2)
	if !bytes.Equal(b1.Bytes(), b2.Bytes()) {
		t.Error()
	}

}

func BenchmarkSSE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		png.Encode(ioutil.Discard, img1)
	}
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		png.Encode(ioutil.Discard, img2)
	}
}
