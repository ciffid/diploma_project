package ui

import (
	"dota3/assets"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
)

type Label struct {
	x, y, size    float64
	px, py, psize int
	textColor     color.Color
	str           *string
	text          *etxt.Renderer
	font          *etxt.Font
}

func NewLabel(x, y, size float64, textColor color.Color, str *string) *Label {
	return &Label{
		x:         x,
		y:         y,
		size:      size,
		psize:     1,
		textColor: textColor,
		str:       str,
		text:      etxt.NewStdRenderer(),
		font:      assets.Fonts["start"],
	}
}

func (l *Label) Update(screen image.Rectangle) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	l.px, l.py = int(l.x*sw), int(l.y*sh)
	l.psize = int(l.size * math.Min(sw, sh))
}

func (l *Label) Draw(screen *ebiten.Image) {
	l.text.SetTarget(screen)
	l.text.SetColor(l.textColor)
	l.text.SetFont(l.font)
	l.text.SetSizePx(l.psize)
	l.text.SetAlign(etxt.YCenter, etxt.XCenter)
	l.text.Draw(*l.str, l.px, l.py)
}
