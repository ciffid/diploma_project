package display

import (
	"dota3/assets"
	"dota3/data"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Camera struct {
	X, Y           float64
	Scale          float64
	screen, target data.Bounder
}

func NewCamera(screen, target data.Bounder) *Camera {
	return &Camera{
		target: target,
		screen: screen,
		Scale:  2,
	}
}

func (c *Camera) Focus() {
	screen := c.screen.Bounds()
	target := c.target.Bounds()
	sw, sh := screen.W, screen.H
	tw, th := target.W, target.H

	scaledTileSize := assets.TileSize * c.Scale
	x, y := target.X*scaledTileSize, target.Y*scaledTileSize

	c.X = x - sw/2 + tw/2*c.Scale
	c.Y = y - sh/2 + th/2*c.Scale
}

func (c *Camera) Zoom() {
	_, delta := ebiten.Wheel()
	if delta > 0 {
		c.Scale *= 1.1
	}
	if delta < 0 {
		c.Scale *= 0.9
	}
	c.Scale = math.Max(0.5, math.Min(5, c.Scale))
}
