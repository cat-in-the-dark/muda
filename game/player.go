package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Player struct {
	pos       *Vector2
	speed     float64
	vp        *Viewport
	state     string
	treasures []*Treasure
	collider  Rect
}

func (p *Player) GetHitRect() Rect {
	return p.collider.Move(p.pos.x, p.pos.y)
}

func NewPlayer(vp *Viewport) *Player {
	return &Player{
		pos:   NewVector2(PlayerStartPosX, PlayerStartPosY),
		speed: 5,
		vp:    vp,
		state: "idle",
		collider: Rect{
			X:      (64 - 40) / 2,
			Y:      0,
			Width:  40,
			Height: 64,
		},
		treasures: make([]*Treasure, 0),
	}
}

func (p *Player) Update() {
	p.state = "idle"
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.pos.y -= p.speed
		p.vp.y -= p.speed
		p.state = "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.pos.y += p.speed
		p.vp.y += p.speed
		p.state = "down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.pos.x -= p.speed
		p.vp.x -= p.speed
		p.state = "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.pos.x += p.speed
		p.vp.x += p.speed
		p.state = "right"
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	var frame *ebiten.Image
	switch p.state {
	case "idle":
		frame = PlayerIdleAnim.GetFrame()
	case "up":
		frame = PlayerUpAnim.GetFrame()
	case "down":
		frame = PlayerDownAnim.GetFrame()
	case "left":
		frame = PlayerRightAnim.GetFrame()
	case "right":
		frame = PlayerRightAnim.GetFrame()
	default:
		log.Panicf("Unknown animation frame %s", p.state)
	}

	opt := &ebiten.DrawImageOptions{}

	x := p.pos.x - p.vp.x
	y := p.pos.y - p.vp.y

	if p.state == "left" {
		opt.GeoM.Scale(-1, 1)
		w, _ := frame.Size()
		opt.GeoM.Translate(float64(w), 0)
	}
	opt.GeoM.Translate(x, y)

	screen.DrawImage(frame, opt)
}

func (p *Player) GetPos() *Vector2 {
	return p.pos
}
