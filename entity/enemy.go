package entity

import (
	"DP/assets"
	"DP/data"
	"DP/display"
	"DP/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Enemy struct {
	X, Y      float64
	Speed     float64
	HitPoints int
	Sprite    *graphics.Frameset
	Direction Direction
	PlayerObj *resolv.Object
}

func NewEnemy(x, y float64) *Enemy {
	return &Enemy{
		X:         x,
		Y:         y,
		Speed:     0.05,
		Sprite:    assets.Images["enemy_look"],
		PlayerObj: resolv.NewObject(x, y, 16, 32),
	}
}

func (e *Enemy) Update() {
	//xMove, yMove := 0., 0.
	//
	//cx, cy := ebiten.CursorPosition()
	//ccx, ccy := (float64(cx) - float64(screen.Bounds().Dx())/2), (float64(cy) - float64(screen.Bounds().Dy())/2)
	//l := math.Sqrt(ccx*ccx + ccy*ccy)
	//angle := math.Atan2(ccy, ccx)
	//
	//if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	//	xMove, yMove = l*math.Cos(angle)/assets.TileSize/camera.Scale, l*math.Sin(angle)/assets.TileSize/camera.Scale
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyW) {
	//	yMove -= e.Speed
	//	e.Direction = DirectionUp
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyS) {
	//	yMove += e.Speed
	//	e.Direction = DirectionDown
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyA) {
	//	xMove -= e.Speed
	//	e.Direction = DirectionLeft
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyD) {
	//	xMove += e.Speed
	//	e.Direction = DirectionRight
	//}
	//e.X += xMove
	//e.Y += yMove
	//
	//playerObj.X += 1
	//playerObj.Y += xMove
	//
	//switch e.Direction {
	//case DirectionLeft:
	//	e.Sprite = assets.Images["player_run_left"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["player_look"]
	//	}
	//case DirectionRight:
	//	e.Sprite = assets.Images["player_run_right"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["player_look"]
	//	}
	//case DirectionDown:
	//	e.Sprite = assets.Images["player_run_down"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["player_look"]
	//	}
	//case DirectionUp:
	//	e.Sprite = assets.Images["player_run_up"]
	//	if xMove == 0 && yMove == 0 {
	//		e.Sprite = assets.Images["player_look"]
	//	}
	//}
	//
	//e.Sprite.Update()
	//e.Box.Update()
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
