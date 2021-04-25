package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Treasure struct {
	id       int32
	texture  *ebiten.Image
	pos      *Vector2
	vp       *Viewport
	trType   int32
	collider Rect
}

func NewTreasure(pos *Vector2, vp *Viewport, ot int32) *Treasure {
	return &Treasure{
		id:      rand.Int31(),
		texture: TreasureTextures[ot],
		pos:     pos,
		vp:      vp,
		trType:  ot,
		collider: Rect{
			X:      0,
			Y:      0,
			Width:  32,
			Height: 32,
		},
	}
}

func (tr *Treasure) Activate() {

}

func (tr *Treasure) Update() {

}

func (tr *Treasure) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	x := tr.pos.x - tr.vp.x
	y := tr.pos.y - tr.vp.y
	opt.GeoM.Translate(x, y)
	screen.DrawImage(tr.texture, opt)
}

func (tr *Treasure) Exit() {

}

func (tr *Treasure) GetHitRect() Rect {
	return tr.collider.Move(tr.pos.x, tr.pos.y)
}

func (tr *Treasure) GetPos() *Vector2 {
	return tr.pos
}
