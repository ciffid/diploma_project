package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Frameset struct {
	Images  []*ebiten.Image
	Ticker  *time.Ticker
	Current int
}

func NewFramesetSingle(frame *ebiten.Image) *Frameset {
	return &Frameset{Ticker: time.NewTicker(1), Images: []*ebiten.Image{frame}}
}

func NewFramesetMulti(speed time.Duration, frames ...*ebiten.Image) *Frameset {
	return &Frameset{Ticker: time.NewTicker(speed), Images: frames}
}

func (s *Frameset) Image() *ebiten.Image {
	return s.Images[s.Current]
}

func (s *Frameset) Update() {
	if len(s.Images) < 2 {
		return
	}
	select {
	case <-s.Ticker.C:
		s.Current = (s.Current + 1) % len(s.Images)
	default:
	}
}
