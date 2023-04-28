package game

import (
	"dota3/data"
	"dota3/display"
	"dota3/levels"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Game struct {
	player *Player
	level  *levels.StartLevel
	camera *display.Camera
	bounds *image.Rectangle
}

func NewGame() *Game {
	g := &Game{
		level:  levels.NewStartLevel(),
		player: NewPlayer(18, 18),
		bounds: &image.Rectangle{},
	}
	bounder := data.NewScreenBounder(g.bounds)
	g.camera = display.NewCamera(bounder, g.player)
	return g
}

func (g *Game) Update() error {
	g.player.Update()
	g.level.Update()
	g.camera.Zoom()
	g.camera.Focus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.level.Draw(screen, g.camera)
	g.player.Draw(screen, g.camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.bounds.Max.X = outsideWidth
	g.bounds.Max.Y = outsideHeight
	return outsideWidth, outsideHeight
}
