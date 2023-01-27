package ui

import "image"

type Bounder interface {
	Bounds() image.Rectangle
}
