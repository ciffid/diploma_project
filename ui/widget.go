package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Widget interface {
	Update(screen image.Rectangle)
	Draw(screen *ebiten.Image)
}
