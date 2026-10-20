package fractx

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func (f *Fractal) AddText(img Image) {

	text := "size=" + f.Size.String() + "  bounds=" + f.Bounds.String()

	bg := color.NRGBA{0, 0, 0, 150}
	fg := color.RGBA{220, 200, 140, 255}

	const (
		margin = 10
		pad    = 4
	)
	face := basicfont.Face7x13
	m := face.Metrics()
	fw := font.MeasureString(face, text).Ceil()
	fh := (m.Ascent + m.Descent).Ceil()

	box := image.Rect(
		margin, margin,
		margin+fw+2*pad, margin+fh+2*pad,
	)
	draw.Draw(img, box, image.NewUniform(bg), image.Point{}, draw.Over)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fg),
		Face: face,
		Dot:  fixed.P(box.Min.X+pad, box.Max.Y-pad-m.Descent.Ceil()),
	}
	d.DrawString(text)
}
