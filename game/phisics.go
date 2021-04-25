package sszb

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r Rect) Move(x, y float64) Rect {
	return Rect{
		X:      r.X + x,
		Y:      r.Y + y,
		Width:  r.Width,
		Height: r.Height,
	}
}

type HitBox interface {
	GetHitRect() Rect
}

// Collide check AABB collisions between two HitBox
func Collide(a, b HitBox) bool {
	rect1 := a.GetHitRect()
	rect2 := b.GetHitRect()

	return rect1.X < rect2.X+rect2.Width &&
		rect1.X+rect1.Width > rect2.X &&
		rect1.Y < rect2.Y+rect2.Height &&
		rect1.Y+rect1.Height > rect2.Y
}
