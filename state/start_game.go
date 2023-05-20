package state

import (
	"DP/display"
	"DP/entity"
	"DP/levels"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type StartGame struct {
	set    SetState
	player *entity.Player
	level  *levels.StartLevel
	camera *display.Camera
	bounds *image.Rectangle
}

func NewStartGame(set SetState) *StartGame {
	s := &StartGame{
		set:    set,
		level:  levels.NewStartLevel(),
		player: entity.NewPlayer(18, 18),
	}
	s.camera = display.NewCamera(s.player)
	return s
}

func (s *StartGame) Update(screen image.Rectangle) error {
	s.player.Update()
	s.level.Update()
	s.camera.Zoom()
	s.camera.Focus(screen)
	return nil
}

func (s *StartGame) Draw(screen *ebiten.Image) {
	s.level.Draw(screen, s.camera)
	s.player.Draw(screen, s.camera)
}

func (s *StartGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s.bounds.Max.X = outsideWidth
	s.bounds.Max.Y = outsideHeight
	return outsideWidth, outsideHeight
}
