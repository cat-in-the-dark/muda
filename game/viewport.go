package sszb

type Viewport struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func NewViewport() *Viewport {
	return &Viewport{
		x:      PlayerStartPosX - ScreenWidth/2,
		y:      PlayerStartPosY - ScreenHeight/2,
		width:  ScreenWidth,
		height: ScreenHeight,
	}
}

func (vp *Viewport) Move(x float64, y float64) {
	vp.x += x
	vp.y += y
}

func (vp *Viewport) Position() (float64, float64) {
	return vp.x, vp.y
}
