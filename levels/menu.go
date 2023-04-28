package levels

import (
	"dota3/assets"
	"dota3/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Menu struct {
	widgets []ui.Widget
}

func NewMenu() *Menu {
	m := &Menu{}
	m.widgets = []ui.Widget{
		//ui.NewBackground(assets.Images["background"], 1, 1),
		ui.NewButton(0, 0, 1, 1,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {},
		),
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
