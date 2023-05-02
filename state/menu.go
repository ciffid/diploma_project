package state

import (
	"dota3/assets"
	"dota3/ui"
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
	textStart := "Начать игру"
	m.widgets = []ui.Widget{
		ui.NewButton(0, 0, 1, 1,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				m.set(NewStartGame(m.set))
			},
		),
		ui.NewLabel(0.5, 0.35, 0.1, colornames.Red, &textStart),
	}
	return m
}

func (g *Menu) Update(screen image.Rectangle) error {
	for _, widget := range g.widgets {
		widget.Update(screen)
	}
	return nil
}

func (g *Menu) Draw(screen *ebiten.Image) {
	for _, widget := range g.widgets {
		widget.Draw(screen)
	}
}
