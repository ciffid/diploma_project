package state

import (
	"game/assets"
	"game/ui"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Fail struct {
	widgets []ui.Widget
	set     SetState
}

func NewFail(set SetState) *Fail {
	s := &Fail{set: set}
	textStart := "Поражение("
	s.widgets = []ui.Widget{
		ui.NewButton(
			0, 0, 1, 1,
			assets.Images["background"],
			assets.Images["background"],
			func() {
				s.set(NewStart(s.set))
			},
		),
		ui.NewLabel(0.5, 0.35, 0.1, colornames.Red, &textStart),
	}
	return s
}

func (g *Fail) Update(screen image.Rectangle) error {
	for _, widget := range g.widgets {
		widget.Update(screen)
	}
	return nil
}

func (g *Fail) Draw(screen *ebiten.Image) {
	for _, widget := range g.widgets {
		widget.Draw(screen)
	}
}
