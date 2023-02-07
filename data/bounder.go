package data

import "image"

type Bounder interface {
	Bounds() image.Rectangle
}

type ScreenBounder struct {
	rect *image.Rectangle
}

func NewScreenBounder(rect *image.Rectangle) *ScreenBounder {
	return &ScreenBounder{
		rect: rect,
	}
}

func (s *ScreenBounder) Bounds() image.Rectangle {
	return *s.rect
}
