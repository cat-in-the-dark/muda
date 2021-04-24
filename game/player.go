package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x       float64
	y       float64
	speed   float64
	texture *ebiten.Image
	vp      *Viewport
}

func NewPlayer(vp *Viewport) *Player {
	return &Player{
		x:       PlayerStartPosX,
		y:       PlayerStartPosY,
		speed:   5,
		texture: PlayerTexture,
		vp:      vp,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= p.speed
		p.vp.y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += p.speed
		p.vp.y += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= p.speed
		p.vp.x -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += p.speed
		p.vp.x += p.speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	x := p.x - p.vp.x
	y := p.y - p.vp.y

	opt.GeoM.Scale(4, 4)
	opt.GeoM.Translate(x, y)
	screen.DrawImage(p.texture, opt)
}
