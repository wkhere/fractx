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
	for _, cmd := range []*exec.Cmd{
		exec.Command("fractx", "-f", "mandelbrot_bw.png", "-bw"),
		exec.Command("convert", "mandelbrot_bw.png", "mandelbrot_bw.ppm"),
		exec.Command("fractx", "-f", "mandelbrot_gs.png"),
		exec.Command("convert", "mandelbrot_gs.png", "mandelbrot_gs.ppm"),
		exec.Command("md5sum", "-c", "MD5"),
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
