package ui

import (
	"DP/graphics"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	backgroundImage *graphics.Frameset
	x, y, w, h      float64
	px, py, pw, ph  float64
}

func NewBackground(x, y, w, h float64, backgroundImage *graphics.Frameset) *Background {
	return &Background{
		backgroundImage: backgroundImage,
		x:               x,
		y:               y,
		w:               w,
		h:               h,
	}
}

func (b *Background) Update(screen image.Rectangle) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	b.px, b.py, b.pw, b.ph = b.x*sw, b.y*sh, b.w*sw, b.h*sh
	b.backgroundImage.Update()
}

func (b *Background) Draw(screen *ebiten.Image) {
	iw, ih := float64(b.backgroundImage.Image().Bounds().Dx()), float64(b.backgroundImage.Image().Bounds().Dy())
	percentIW, percentIH := b.pw/iw, b.ph/ih

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(percentIW, percentIH)
	op.GeoM.Translate(b.px, b.py)
	screen.DrawImage(b.backgroundImage.Image(), op)
}
