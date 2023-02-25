package levels

import (
	"dota3/assets"
	"dota3/display"
	"dota3/tilemap"
	"dota3/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type StartLevel struct {
	location *ui.Background
	tilemap  *tilemap.Tilemap
}

func NewStartLevel() *StartLevel {
	s := &StartLevel{
		location: ui.NewBackground(assets.Images["start-location"], 1.2, 1.2),
		tilemap:  tilemap.NewTilemap(36, 12),
	}
	return s
}

func (s *StartLevel) Update() {
}

func (s *StartLevel) Draw(screen *ebiten.Image, camera *display.Camera) {
	s.location.Draw(screen, camera)
	s.tilemap.Draw(screen, camera)
}
