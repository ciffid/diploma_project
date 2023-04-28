package data

type Bounds struct {
	X, Y, W, H float64
}

func NewBounds(x, y, w, h float64) Bounds {
	return Bounds{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}
