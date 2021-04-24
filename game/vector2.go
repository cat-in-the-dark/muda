package sszb

type Vector2 struct {
	x float64
	y float64
}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{x, y}
}
