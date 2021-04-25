package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Treasure struct {
	texture *ebiten.Image
	pos     *Vector2
	vp      *Viewport
}

func NewTreasure(pos *Vector2, vp *Viewport) *Treasure {
	return &Treasure{
		texture: TreasureTexture,
		pos:     pos,
		vp:      vp,
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
