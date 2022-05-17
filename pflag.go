package fractx

import (
	"fmt"
	"strconv"
	"strings"
)

// support for pflag.Value

func (s *Size) Set(input string) (err error) {
	cc := strings.Split(input, ",")
	if len(cc) != 2 {
		return fmt.Errorf("expected width,height")
	}
	s.W, err = strconv.Atoi(strings.TrimSpace(cc[0]))
	if err != nil {
		return err
	}
	if s.W <= 0 {
		return fmt.Errorf(`parsed "%d": width <= 0`, s.W)
	}
	s.H, err = strconv.Atoi(strings.TrimSpace(cc[1]))
	if err != nil {
		return err
	}
	if s.H <= 0 {
		return fmt.Errorf(`parsed "%d": height <= 0`, s.H)
	}
	return nil
}

func (s *Size) String() string {
	return fmt.Sprintf("%d,%d", s.W, s.H)
}

func (*Size) Type() string {
	return "width,height"
}

func (r *Rect) Set(input string) (err error) {
	cc := strings.Split(input, ",")
	if len(cc) != 4 {
		return fmt.Errorf("expected x0,y0,x1,y1")
	}
	r.X0, err = strconv.ParseFloat(strings.TrimSpace(cc[0]), 64)
	if err != nil {
		return err
	}
	r.Y0, err = strconv.ParseFloat(strings.TrimSpace(cc[1]), 64)
	if err != nil {
		return err
	}
	r.X1, err = strconv.ParseFloat(strings.TrimSpace(cc[2]), 64)
	if err != nil {
		return err
	}
	r.Y1, err = strconv.ParseFloat(strings.TrimSpace(cc[3]), 64)
	if err != nil {
		return err
	}
	return nil
}

func (r *Rect) String() string {
	return fmt.Sprintf("%g,%g,%g,%g", r.X0, r.Y0, r.X1, r.Y1)
}

func (*Rect) Type() string {
	return "x0,y0,x1,y1"
}
