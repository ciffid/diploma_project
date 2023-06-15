package state

import (
	"DP/assets"
	"DP/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image"
)

type Menu struct {
	widgets []ui.Widget
	set     SetState
}

func NewMenu(set SetState) *Menu {
	m := &Menu{set: set}
	textName := "game name"
	textPlay := "play"
	textSettings := "settings"
	textExit := "exit"
	m.widgets = []ui.Widget{
		ui.NewBackground(0, 0, 1, 1,
			assets.Images["background"]),
		ui.NewButton(0.3, 0.4, 0.4, 0.1,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				m.set(NewStartGame(m.set))
			},
		),
		ui.NewButton(0.3, 0.55, 0.4, 0.1,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				m.set(NewSettings(m.set))
			},
		),
		ui.NewButton(0.3, 0.7, 0.4, 0.1,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
			},
		),
		ui.NewLabel(0.5, 0.15, 0.12, colornames.Black, &textName),
		ui.NewLabel(0.5, 0.45, 0.08, colornames.Black, &textPlay),
		ui.NewLabel(0.5, 0.6, 0.06, colornames.Black, &textSettings),
		ui.NewLabel(0.5, 0.75, 0.08, colornames.Black, &textExit),
	}
	return m
}

func (m *Menu) Update(screen image.Rectangle) error {
	for _, widget := range m.widgets {
		widget.Update(screen)
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for _, widget := range m.widgets {
		widget.Draw(screen)
	}
}
