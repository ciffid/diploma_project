package game

import (
	"dota3/assets"
	"dota3/data"
	"dota3/display"
	"github.com/hajimehoshi/ebiten/v2"
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
		Speed:     0.05,
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

func (p *Player) Draw(screen *ebiten.Image, camera *display.Camera) {
	op := &ebiten.DrawImageOptions{}
	if p.Direction == DirectionRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(p.Sprite.Bounds().Dx()), 0)
	}
	scaledTileSize := assets.TileSize * camera.Scale
	op.GeoM.Scale(camera.Scale, camera.Scale)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Translate(p.X*scaledTileSize, p.Y*scaledTileSize)
	screen.DrawImage(p.Sprite, op)
}

func (p *Player) Bounds() data.Bounds {
	return data.NewBounds(p.X, p.Y, float64(p.Sprite.Bounds().Dx()), float64(p.Sprite.Bounds().Dy()))
}
