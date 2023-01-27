package ui

import (
	"dota3/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Background struct {
	locationImage  *ebiten.Image
	scaleW, scaleH float64
}

func NewBackground(locationImage *ebiten.Image, scaleW, scaleH float64) *Background {
	return &Background{
		locationImage: locationImage,
		scaleW:        scaleW,
		scaleH:        scaleH,
	}
}

func (l *Background) Update(screen image.Rectangle) {

}

func (l *Background) Draw(screen *ebiten.Image, camera *game.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-)
	op.GeoM.Scale(l.scaleW, l.scaleH)
	screen.DrawImage(l.locationImage, op)
}

func (l *Background) Bounds() image.Rectangle {
	w, h := float64(l.locationImage.Bounds().Dx()), float64(l.locationImage.Bounds().Dy())
	return image.Rect(0, 0, int(w*l.scaleW), int(h*l.scaleH))
}
