package main

import (
	"embed"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/wkhere/fractx"
)

//go:embed favicon.ico
var fs embed.FS

type server struct {
	fractal  *fractx.Fractal
	origSize fractx.Size
	newImage fractx.ImageBuilder
	addr     string
}

func (s *server) serve() error {
	u, err := browserURL(s.addr)
	if err != nil {
		return err
	}

	files := http.FileServer(http.FS(fs))

	http.Handle("/", redirect("/f", 303))
	http.Handle("/f", http.HandlerFunc(fractalHandler(s)))
	http.Handle("/favicon.ico", files)

	fmt.Println("visit", u)
	return http.ListenAndServe(s.addr, nil)
}

func browserURL(addr string) (string, error) {
	u, err := url.Parse("http://" + addr)
	if err != nil {
		return "", err
	}
	if u.Hostname() == "" {
		return "http://localhost:" + u.Port(), nil
	}
	return u.String(), nil
}

func fractalHandler(s *server) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		// todo: handle err

		var pixw, pixh int
		var ok bool
		size := s.origSize

		pixw, ok, err = sizeParam(req.Form.Get("w"))
		if err != nil {
			log.Printf("invalid width: %s; using default=%d", err, size.W)
		}
		if ok {
			size.W = pixw
		}

		pixh, ok, err = sizeParam(req.Form.Get("h"))
		if err != nil {
			log.Printf("invalid height: %s; using default=%d", err, size.H)
		}
		if ok {
			size.H = pixh
		}

		w.Header().Set("Content-Type", "image/png")

		s.fractal.Size = size
		img := s.newImage(s.fractal)
		s.fractal.Fill(img)

		err = png.Encode(w, img)
		// todo: handle err
	}
}

func sizeParam(val string) (x int, ok bool, err error) {
	if val == "" {
		return 0, false, nil
	}
	x, err = strconv.Atoi(val)
	if err != nil {
		return 0, false, err
	}
	if err == nil && x <= 0 {
		return 0, false, fmt.Errorf("%q is not a positive int", val)
	}
	return x, true, nil
}

func redirect(path string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, code)
	}
}
