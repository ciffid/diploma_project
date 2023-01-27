package game

import "dota3/ui"

type Camera struct {
	X, Y   float64
	target ui.Bounder
}

func NewCamera(target ui.Bounder) *Camera {
	return &Camera{
		target: target,
	}
}

func (c *Camera) Focus() {
	rect := c.target.Bounds()
	c.X, c.Y = float64(rect.Min.X+rect.Dx()/2), float64(rect.Min.Y+rect.Dy()/2)
}
