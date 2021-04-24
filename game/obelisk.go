package sszb

import "github.com/hajimehoshi/ebiten/v2"

// Obelisk is an Object where the Player supposed to put a Treasure.
type Obelisk struct {
	x          float64
	y          float64
	texture    *ebiten.Image
	nTreasures int
	vp         *Viewport
}

func NewObelisk(vp *Viewport) *Obelisk {
	return &Obelisk{
		x:          0,
		y:          0,
		texture:    ObeliskTexture,
		nTreasures: 0,
		vp:         vp,
	}
}

func (o *Obelisk) Update() {

}

func (o *Obelisk) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	x := o.x - o.vp.x
	y := o.y - o.vp.y
	opt.GeoM.Translate(x, y)
	screen.DrawImage(o.texture, opt)
}
