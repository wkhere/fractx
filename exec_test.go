package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func init() {
	cwd, _ := os.Getwd()
	os.Setenv("PATH", cwd+string(os.PathListSeparator)+os.Getenv("PATH"))
	os.Chdir("testdata")
}

func TestExec(t *testing.T) {
	c := exec.Command
	for _, cmd := range []*exec.Cmd{
		c("fractx", "-o", "mandelbrot_bw.png", "-color=bw"),
		c("convert", "mandelbrot_bw.png", "mandelbrot_bw.ppm"),
		c("fractx", "-o", "mandelbrot_gray.png"),
		c("convert", "mandelbrot_gray.png", "mandelbrot_gray.ppm"),
		c("md5sum", "-c", "MD5"),
	} {
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Errorf("%q error:\n%s\n%s",
				strings.Join(cmd.Args, " "), err, out)
			return
		}
	}
}

func BenchmarkExecBW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("fractx", "-f", "mandelbrot.png", "-bw").Run()
	}
}

func BenchmarkExecGS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("fractx", "-f", "mandelbrot.png").Run()
	}
}
