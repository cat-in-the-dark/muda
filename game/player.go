package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x       float64
	y       float64
	speed   float64
	vp      *Viewport
	state string
}

func NewPlayer(vp *Viewport) *Player {
	return &Player{
		x:         PlayerStartPosX,
		y:         PlayerStartPosY,
		speed:     5,
		vp:        vp,
		state:     "idle",
	}
}

func (p *Player) Update() {
	p.state = "idle"
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= p.speed
		p.vp.y -= p.speed
		p.state = "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += p.speed
		p.vp.y += p.speed
		p.state = "down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= p.speed
		p.vp.x -= p.speed
		p.state = "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += p.speed
		p.vp.x += p.speed
		p.state = "right"
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	x := p.x - p.vp.x
	y := p.y - p.vp.y

	opt.GeoM.Translate(x, y)

	switch p.state {
	case "idle":
		screen.DrawImage(PlayerIdleAnim.GetFrame(), opt)
	case "up":
		screen.DrawImage(PlayerUpAnim.GetFrame(), opt)
	case "down":
		screen.DrawImage(PlayerDownAnim.GetFrame(), opt)
	case "left":
		screen.DrawImage(PlayerLeftAnim.GetFrame(), opt)
	case "right":
		screen.DrawImage(PlayerRightAnim.GetFrame(), opt)
	}

}
