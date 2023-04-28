package data

import "image"

type ScreenBounder struct {
	rect *image.Rectangle
}

func NewScreenBounder(rect *image.Rectangle) *ScreenBounder {
	return &ScreenBounder{
		rect: rect,
	}
}

func (s *ScreenBounder) Bounds() Bounds {
	return NewBounds(
		float64(s.rect.Min.X),
		float64(s.rect.Min.Y),
		float64(s.rect.Dx()),
		float64(s.rect.Dy()),
	)
}
