package display

import (
	"dota3/data"
)

type Camera struct {
	X, Y           float64
	screen, target data.Bounder
}

func NewCamera(screen, target data.Bounder) *Camera {
	return &Camera{
		target: target,
		screen: screen,
	}
}

func (c *Camera) Focus() {
	target := c.target.Bounds()
	screen := c.screen.Bounds()
	tw, th := float64(target.Dx()), float64(target.Dy())
	sw, sh := float64(screen.Dx()), float64(screen.Dy())
	x, y := float64(target.Min.X), float64(target.Min.Y)
	c.X = x - (sw-tw)/2
	c.Y = y - (sh-th)/2
}
