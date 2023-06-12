package state

import (
	"DP/assets"
	"DP/display"
	"DP/entity"
	"DP/tilemap"
	"DP/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"golang.org/x/image/colornames"
	"image"
)

type StartGame struct {
	set     SetState
	player  *entity.Player
	tilemap *tilemap.Tilemap
	camera  *display.Camera
	bounds  *image.Rectangle
	space   *resolv.Space
	widgets []ui.Widget
}

func NewStartGame(set SetState) *StartGame {
	s := &StartGame{
		set:   set,
		space: resolv.NewSpace(assets.TileMapSize, assets.TileMapSize, 1, 1),
	}
	textName := "menu"
	s.tilemap = tilemap.NewTilemap(assets.TileMapSize, assets.TileMapSize, s.space)
	s.player = entity.NewPlayer(18, 18, s.space)
	s.camera = display.NewCamera(s.player)
	s.widgets = []ui.Widget{
		ui.NewButton(0.01, 0.025, 0.08, 0.05,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewMenu(s.set))
			},
		),
		ui.NewLabel(0.05, 0.05, 0.03, colornames.Black, &textName),
	}
	return s
}

func (s *StartGame) Update(screen image.Rectangle) error {
	s.player.Update(screen, s.camera)
	s.camera.Zoom()
	s.camera.Focus(screen)
	for _, widget := range s.widgets {
		widget.Update(screen)
	}
	return nil
}

func (s *StartGame) Draw(screen *ebiten.Image) {
	s.tilemap.Draw(screen, s.camera)
	s.player.Draw(screen, s.camera)
	for _, widget := range s.widgets {
		widget.Draw(screen)
	}
}

func (s *StartGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s.bounds.Max.X = outsideWidth
	s.bounds.Max.Y = outsideHeight
	return outsideWidth, outsideHeight
}
