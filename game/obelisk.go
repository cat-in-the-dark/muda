package sszb

import "github.com/hajimehoshi/ebiten/v2"

// Obelisk is an Object where the Player supposed to put a Treasure.
type Obelisk struct {
	pos          *Vector2
	texture      *ebiten.Image
	treasureType int32
	vp           *Viewport
}

func NewObelisk(pos *Vector2, vp *Viewport, tt int32) *Obelisk {
	return &Obelisk{
		pos:          pos,
		texture:      ObeliskTexture,
		treasureType: tt,
		vp:           vp,
	}
}

func (o *Obelisk) Update() {

}

func (o *Obelisk) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	x := o.pos.x - o.vp.x
	y := o.pos.y - o.vp.y
	opt.GeoM.Translate(x, y)
	screen.DrawImage(o.texture, opt)
}
