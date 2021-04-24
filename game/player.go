package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x       float64
	y       float64
	speed   float64
	texture *ebiten.Image
}

func NewPlayer() *Player {
	return &Player{
		x: 0,
		y: 0,
		speed: 1,
		texture: PlayerTexture,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += p.speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.x, p.y)
	opt.GeoM.Scale(4, 4)
	screen.DrawImage(p.texture, opt)
}
