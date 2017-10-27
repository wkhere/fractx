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
		exec.Command("fractx", "-f", "mandelbrot.png"),
		exec.Command("convert", "mandelbrot.png", "mandelbrot.ppm"),
		exec.Command("md5sum", "-c", "MD5"),
	} {
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Errorf("%q error:\n%s\n%s",
				strings.Join(cmd.Args, " "), err, out)
			return
		}
	}
}

func BenchmarkExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("fractx", "-f", "mandelbrot.png").Run()
	}
}
