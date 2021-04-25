package sszb

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	GetPos() *Vector2
	Draw(screen *ebiten.Image)
}
