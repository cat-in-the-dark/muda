package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	vp     *Viewport
	player *Player
}

func NewGameScene() *GameScene {
	vp := NewViewport()
	return &GameScene{
		vp:     vp,
		player: NewPlayer(vp),
	}
}

func (g *GameScene) Activate() {

}

func (g *GameScene) Update() {
	g.player.Update()
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-g.vp.x, -g.vp.y)
	screen.DrawImage(BG, opt)

	g.player.Draw(screen)
}

func (g *GameScene) Exit() {

}
