package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tree struct {
	texture *ebiten.Image
	pos     *Vector2
	vp      *Viewport
}

func NewTree(pos *Vector2, vp *Viewport) *Tree {
	return &Tree{
		texture: TreeTexture,
		pos:     pos,
		vp:      vp,
	}
}

func (tr *Tree) Activate() {
}

func (tr *Tree) Update() {
}

func (tr *Tree) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	x := tr.pos.x - tr.vp.x
	y := tr.pos.y - tr.vp.y
	opt.GeoM.Translate(x, y)
	screen.DrawImage(tr.texture, opt)
}

func (tr *Tree) Exit() {
}

func (tr *Tree) GetPos() *Vector2 {
	return tr.pos
}
