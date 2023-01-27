package game

import (
	"dota3/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
)

type Player struct {
	X, Y      float64
	Speed     float64
	Direction Direction
	Sprite    *ebiten.Image
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:         x,
		Y:         y,
		Speed:     2,
		Direction: DirectionLeft,
		Sprite:    assets.Images["player"],
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.Speed
		p.Direction = DirectionLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.Speed
		p.Direction = DirectionRight
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if p.Direction == DirectionRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(p.Sprite.Bounds().Dx()), 0)
	}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, op)
}

func (p *Player) Bounds() image.Rectangle {
	return image.Rect(int(p.X), int(p.Y), int(p.X)+p.Sprite.Bounds().Dx(), int(p.Y)+p.Sprite.Bounds().Dy())
}
