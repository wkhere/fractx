package main

import (
	"embed"
	"fmt"
	"image/png"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/wkhere/fractx"
)

//go:embed index.html favicon.ico
var fs embed.FS

type server struct {
	fractal0 fractx.Fractal
	newImage fractx.ImageBuilder
	addr     string
	text     bool
}

func (s *server) serve() error {
	lst, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	files := http.FileServer(http.FS(fs))

	http.Handle("/", files)
	http.Handle("/f", http.HandlerFunc(fractalHandler(s)))

	fmt.Println("visit", urlFromAddr(lst.Addr()))
	return (&http.Server{}).Serve(lst)
}

func urlFromAddr(addr net.Addr) *url.URL {
	u, err := url.Parse("http://" + addr.String())
	if err != nil {
		panic(err) // shouldn't happen as the addr is already used by the listener
	}
	if u.Hostname() == "" {
		u.Host = "localhost:" + u.Port()
	}
	return u
}

func fractalHandler(s *server) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), 400)
			log.Println("parse form:", err)
			return
		}

		f := s.fractal0

		x, ok, err := sizeParam(req.Form.Get("w"))
		if err != nil {
			log.Printf("invalid width: %s; using default=%d", err, f.Size.W)
		}
		if ok {
			f.Size.W = x
		}

		x, ok, err = sizeParam(req.Form.Get("h"))
		if err != nil {
			log.Printf("invalid height: %s; using default=%d", err, f.Size.H)
		}
		if ok {
			f.Size.H = x
		}

		if s := req.Form.Get("size"); s != "" {
			var nsize fractx.Size
			err := nsize.Set(s)
			if err != nil {
				log.Printf("invalid size: %s; using default=%s", err, f.Size)
			} else {
				f.Size = nsize
			}
		}

		if s := req.Form.Get("bounds"); s != "" {
			var nbounds fractx.Rect
			err := nbounds.Set(s)
			if err != nil {
				log.Printf("invalid bounds: %s; using default=%s", err, f.Bounds)
			} else {
				f.Bounds = nbounds
			}
		}

		nmaxi, ok, err := uintParam(req.Form.Get("maxi"))
		if err != nil {
			log.Printf("invalid maxi: %s; using default=%d", err, f.MaxI)
		}
		if ok {
			f.MaxI = nmaxi
		}

		w.Header().Set("Content-Type", "image/png")

		img := s.newImage(&f)
		f.Fill(img)
		if s.text {
			f.AddText(img)
		}

		err = png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println("encode png:", err)
			return
		}
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
	if x <= 0 {
		return 0, false, fmt.Errorf("%q is not a positive int", val)
	}
	return x, true, nil
}

func uintParam(val string) (_ uint, ok bool, err error) {
	if val == "" {
		return 0, false, nil
	}
	x, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return 0, false, err
	}
	return uint(x), true, nil
}

func redirect(path string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, code)
	}
}
