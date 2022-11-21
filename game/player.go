package game

import (
	"dota3/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y float64
}

func NewPLayer(x, y float64) *Player {
	return &Player{
		X: x,
		Y: y,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += 5
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(assets.Images["player"], op)
}
