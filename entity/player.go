package entity

import (
	"DP/assets"
	"DP/data"
	"DP/display"
	"DP/graphics"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
	"image"
	"math"
)

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionUp
	DirectionDown
)

type Player struct {
	X, Y      float64
	Speed     float64
	HitPoints int
	Sprite    *graphics.Frameset
	Direction Direction
	Box       *resolv.Object
}

func NewPlayer(x, y float64, space *resolv.Space) *Player {
	p := &Player{
		X:         x,
		Y:         y,
		Speed:     0.05,
		Direction: DirectionRight,
		Sprite:    assets.Images["player_run_right"],
	}
	p.Box = resolv.NewObject(x, y, 0.5, 0.6)
	space.Add(p.Box)
	return p
}

func (p *Player) Update(screen image.Rectangle, camera *display.Camera) {
	xMove, yMove := 0., 0.

	cx, cy := ebiten.CursorPosition()
	ccx, ccy := float64(cx)-float64(screen.Bounds().Dx())/2, float64(cy)-float64(screen.Bounds().Dy())/2
	l := math.Sqrt(ccx*ccx + ccy*ccy)
	angle := math.Atan2(ccy, ccx)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		xMove, yMove = l*math.Cos(angle)/assets.TileSize/camera.Scale, l*math.Sin(angle)/assets.TileSize/camera.Scale
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		yMove -= p.Speed
		p.Direction = DirectionUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		yMove += p.Speed
		p.Direction = DirectionDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		xMove -= p.Speed
		p.Direction = DirectionLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		xMove += p.Speed
		p.Direction = DirectionRight
	}

	switch p.Direction {
	case DirectionLeft:
		p.Sprite = assets.Images["player_run_left"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look"]
		}
	case DirectionRight:
		p.Sprite = assets.Images["player_run_right"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look"]
		}
	case DirectionDown:
		p.Sprite = assets.Images["player_run_down"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look"]
		}
	case DirectionUp:
		p.Sprite = assets.Images["player_run_up"]
		if xMove == 0 && yMove == 0 {
			p.Sprite = assets.Images["player_look"]
		}
	}
	if collision := p.Box.Check(xMove, yMove); collision != nil {
		contact := collision.ContactWithObject(collision.Objects[0])
		xMove, yMove = 0, 0
		fmt.Println(contact.X(), contact.Y())
	}
	p.move(xMove, yMove)
	p.Sprite.Update()
	p.Box.Update()
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

func (p *Player) move(dx, dy float64) {
	p.X += dx
	p.Y += dy
	p.Box.X += dx
	p.Box.Y += dy
}
