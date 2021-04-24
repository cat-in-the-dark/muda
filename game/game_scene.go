package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct{
	player* Player
}

func NewGameScene() *GameScene {
	return &GameScene{
		player: NewPlayer(),
	}
}

func (g *GameScene) Activate() {

}

func (g *GameScene) Update() {
	g.player.Update()
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *GameScene) Exit() {

}
