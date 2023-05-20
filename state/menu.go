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
	textStart := "Начать игру"
	m.widgets = []ui.Widget{
		ui.NewBackground(0, 0, 1, 1,
			assets.Images["background"]),
		ui.NewButton(0.5-0.5/2, 0.35-0.2/2, 0.5, 0.2,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				m.set(NewStartGame(m.set))
			},
		),
		ui.NewLabel(0.5, 0.35, 0.1, colornames.Blue, &textStart),
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
