package state

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Update(screen image.Rectangle) error
	Draw(screen *ebiten.Image)
}

type SetState func(State)
