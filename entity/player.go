package entity

import (
	"DP/assets"
	"DP/data"
	"DP/display"
	"DP/graphics"
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
	HitPoints int
	Sprite    *graphics.Frameset
	Direction Direction
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:         x,
		Y:         y,
		Speed:     0.05,
		Direction: DirectionRight,
		Sprite:    assets.Images["player_run_right"],
	}
}

func (p *Player) Update() {
	xMove, yMove := 0., 0.
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		yMove -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		xMove -= p.Speed
		p.Direction = DirectionLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		yMove += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		xMove += p.Speed
		p.Direction = DirectionRight
	}
	p.X += xMove
	p.Y += yMove
	switch p.Direction {
	case DirectionLeft:
		p.Sprite = assets.Images["player_run_left"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look_left"]
		}
	case DirectionRight:
		p.Sprite = assets.Images["player_run_right"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look_right"]
		}
	}
	p.Sprite.Update()

}

func (p *Player) Draw(screen *ebiten.Image, camera *display.Camera) {
	op := &ebiten.DrawImageOptions{}
	scaledTileSize := assets.TileSize * camera.Scale
	op.GeoM.Scale(camera.Scale, camera.Scale)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Translate(p.X*scaledTileSize, p.Y*scaledTileSize)
	screen.DrawImage(p.Sprite.Image(), op)
}

func (p *Player) Bounds() data.Bounds {
	return data.NewBounds(p.X, p.Y, float64(p.Sprite.Image().Bounds().Dx()), float64(p.Sprite.Image().Bounds().Dy()))
}
