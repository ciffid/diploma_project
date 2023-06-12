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
		ui.NewButton(0.5-0.5/2, 0.45-0.2/2, 0.5, 0.15,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				m.set(NewStartGame(m.set))
			},
		),
		ui.NewButton(0.5-0.5/2, 0.65-0.2/2, 0.5, 0.15,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
			},
		),
		ui.NewButton(0.5-0.5/2, 0.85-0.2/2, 0.5, 0.15,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
			},
		),
		ui.NewLabel(0.5, 0.13, 0.16, colornames.Black, &textName),
		ui.NewLabel(0.5, 0.43, 0.1, colornames.Black, &textPlay),
		ui.NewLabel(0.5, 0.63, 0.07, colornames.Black, &textSettings),
		ui.NewLabel(0.5, 0.83, 0.09, colornames.Black, &textExit),
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