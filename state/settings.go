package state

import (
	"DP/assets"
	"DP/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image"
)

type Settings struct {
	widgets []ui.Widget
	set     SetState
}

func NewSettings(set SetState) *Settings {
	s := &Settings{set: set}
	textSettings := "settings"
	textAudio := "audio"
	textGraphics := "graphics"
	textGame := "game"
	textKeyBinder := "key binder"
	textBack := "back"
	s.widgets = []ui.Widget{

		ui.NewBackground(0, 0, 1, 1,
			assets.Images["background"]),
		ui.NewBackground(0, 0, 1, 1,
			assets.Images["settings"]),

		ui.NewButton(0.3, 0.35, 0.4, 0.09,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewStartGame(s.set))
			},
		),
		ui.NewButton(0.3, 0.45, 0.4, 0.09,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewStartGame(s.set))
			},
		),
		ui.NewButton(0.3, 0.55, 0.4, 0.09,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewSettings(s.set))
			},
		),
		ui.NewButton(0.3, 0.65, 0.4, 0.09,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewStartGame(s.set))
			},
		),
		ui.NewButton(0.3, 0.75, 0.4, 0.09,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.set(NewMenu(s.set))
			},
		),
		ui.NewLabel(0.5, 0.15, 0.14, colornames.White, &textSettings),
		ui.NewLabel(0.5, 0.4, 0.08, colornames.Black, &textGame),
		ui.NewLabel(0.5, 0.5, 0.067, colornames.Black, &textAudio),
		ui.NewLabel(0.5, 0.6, 0.06, colornames.Black, &textGraphics),
		ui.NewLabel(0.5, 0.7, 0.047, colornames.Black, &textKeyBinder),
		ui.NewLabel(0.5, 0.8, 0.075, colornames.Black, &textBack),
	}
	return s
}

func (s *Settings) Update(screen image.Rectangle) error {
	for _, widget := range s.widgets {
		widget.Update(screen)
	}
	return nil
}

func (s *Settings) Draw(screen *ebiten.Image) {
	for _, widget := range s.widgets {
		widget.Draw(screen)
	}
}
