package display

import (
	"DP/assets"
	"DP/data"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

type Camera struct {
	X, Y   float64
	Scale  float64
	target data.Bounder
}

func NewCamera(target data.Bounder) *Camera {
	return &Camera{
		target: target,
		Scale:  6,
	}
}

func (c *Camera) Focus(screen image.Rectangle) {
	target := c.target.Bounds()
	sw, sh := float64(screen.Dx()), float64(screen.Dy())
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
	c.Scale = math.Max(4, math.Min(8, c.Scale))
}
