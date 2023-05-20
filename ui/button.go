package ui

import (
	"DP/graphics"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	defaultImage, pressedImage *graphics.Frameset
	buttonPressed              bool
	x, y, w, h                 float64
	px, py, pw, ph             float64
	callback                   func()
}

func NewButton(x, y, w, h float64, defaultImage, pressedImage *graphics.Frameset, callback func()) *Button {
	return &Button{
		defaultImage: defaultImage,
		pressedImage: pressedImage,
		x:            x,
		y:            y,
		w:            w,
		h:            h,
		callback:     callback,
	}
}

func (b *Button) Update(screen image.Rectangle) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	b.px, b.py, b.pw, b.ph = b.x*sw, b.y*sh, b.w*sw, b.h*sh

	cx, cy := ebiten.CursorPosition()
	ccx, ccy := float64(cx), float64(cy)
	if ccx >= b.px && ccx < b.px+b.pw && ccy >= b.py && ccy < b.py+b.ph {
		b.buttonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			b.callback()
		}
	}
	b.defaultImage.Update()
	b.pressedImage.Update()
}

func (b *Button) Draw(screen *ebiten.Image) {
	iw, ih := float64(b.defaultImage.Image().Bounds().Dx()), float64(b.defaultImage.Image().Bounds().Dy())
	percentIW, percentIH := b.pw/iw, b.ph/ih

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(percentIW, percentIH)
	op.GeoM.Translate(b.px, b.py)

	button := b.defaultImage
	if b.buttonPressed {
		button = b.pressedImage
	}
	screen.DrawImage(button.Image(), op)
}
