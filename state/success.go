package state

import (
	"dota3/assets"
	"dota3/ui"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Success struct {
	widgets []ui.Widget
	set     SetState
}

func NewSuccess(set SetState) *Success {
	s := &Success{set: set}
	textStart := "Победа!"
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

func (g *Success) Update(screen image.Rectangle) error {
	for _, widget := range g.widgets {
		widget.Update(screen)
	}
	return nil
}

func (g *Success) Draw(screen *ebiten.Image) {
	for _, widget := range g.widgets {
		widget.Draw(screen)
	}
}
