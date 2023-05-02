package game

import (
	"dota3/state"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State  state.State
	Screen image.Rectangle
}

func NewGame() *Game {
	g := &Game{}
	g.State = state.NewMenu(g.SetState)
	return g
}
func (g *Game) SetState(next state.State) {
	g.State = next
}
func (g *Game) Update() error {
	err := g.State.Update(g.Screen)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.Screen = image.Rect(0, 0, outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
