package sszb

import (
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	GetPos() *Vector2
	Draw(screen *ebiten.Image)
}

type AnimationAt struct {
	animation     *lib.Animation
	pos           *Vector2
	vp            *Viewport
	elapsedFrames int64
	totalFrames   int64
}

func NewAnimationAt(animation *lib.Animation, pos *Vector2, vp *Viewport) *AnimationAt {
	return &AnimationAt{
		animation:     animation,
		pos:           pos,
		vp:            vp,
		elapsedFrames: 0,
		totalFrames:   80,
	}
}

func (a *AnimationAt) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(a.pos.x-a.vp.x, a.pos.y-a.vp.y)
	screen.DrawImage(a.animation.GetFrame(), opts)
	a.elapsedFrames++
}

func (a *AnimationAt) IsFinished() bool {
	return a.elapsedFrames >= a.totalFrames
}
