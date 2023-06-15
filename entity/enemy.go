package entity

import (
	"DP/assets"
	"DP/data"
	"DP/display"
	"DP/graphics"
	"fmt"
	"time"
)

type Enemy struct {
	X, Y      float64
	Speed     float64
	HitPoints int
	Sprite    *graphics.Frameset
	Direction Direction
	Ticker    *time.Ticker
}

func NewEnemy(x, y float64) *Enemy {
	return &Enemy{
		X:         x,
		Y:         y,
		Speed:     0.03,
		Sprite:    assets.Images["enemy_look"],
		Direction: DirectionRight,
		Ticker:    time.NewTicker(time.Second * 3),
	}
}

func (e *Enemy) Update() {

	select {
	case <-e.Ticker.C:
		e.Speed *= -1
		fmt.Println(e.Speed)
	default:
	}

	e.X += e.Speed
	e.Y += 0

	if e.Speed < 0 {
		e.Direction = DirectionLeft
		e.Sprite = assets.Images["enemy_run_left"]
	}
	//if xMove == 0 && yMove == 0 {
	//	e.Sprite = assets.Images["enemy_look"]
	//}
	if e.Speed > 0 {
		e.Direction = DirectionLeft
		e.Sprite = assets.Images["enemy_run_right"]
	}
	//if xMove == 0 && yMove == 0 {
	//	e.Sprite = assets.Images["enemy_look"]
	//}
	//case DirectionDown:
	//	e.Sprite = assets.Images["enemy_run_down"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["enemy_look"]
	//	}
	//case DirectionUp:
	//	e.Sprite = assets.Images["enemy_run_up"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["enemy_look"]
	//	}

	e.Sprite.Update()
}

func (e *Enemy) Draw(screen *ebiten.Image, camera *display.Camera) {
	op := &ebiten.DrawImageOptions{}
	scaledTileSize := assets.TileSize * camera.Scale
	op.GeoM.Scale(camera.Scale, camera.Scale)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Translate(e.X*scaledTileSize, e.Y*scaledTileSize)
	screen.DrawImage(e.Sprite.Image(), op)
}

func (e *Enemy) Bounds() data.Bounds {
	return data.NewBounds(e.X, e.Y, float64(e.Sprite.Image().Bounds().Dx()), float64(e.Sprite.Image().Bounds().Dy()))
}
