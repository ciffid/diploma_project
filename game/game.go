package game

import (
	"dota3/levels"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
	level  *levels.StartLevel
	camera *Camera
}

func NewGame() *Game {
	g := &Game{
		player: NewPlayer(100, 100),
		level:  levels.NewStartLevel(),
	}
	g.camera = NewCamera(g.player)
	return g
}

func (g *Game) Update() error {
	g.player.Update()
	g.level.Update()
	g.camera.Focus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.level.Draw(screen, g.camera)
	g.player.Draw(screen, g.camera)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
