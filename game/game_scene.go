package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	vp           *Viewport
	player       *Player
	obelisk      *Obelisk
	mapGenerator *MapGenerator
	gameMap      *Map
}

func NewGameScene() *GameScene {
	vp := NewViewport()
	return &GameScene{
		vp:           vp,
		player:       NewPlayer(vp),
		obelisk:      NewObelisk(vp),
		mapGenerator: NewMapGenerator(20, vp),
	}
}

func (g *GameScene) Activate() {
	g.gameMap = g.mapGenerator.Generate()
}

func (g *GameScene) Update() {
	g.player.Update()
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(ColorBack)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-g.vp.x, -g.vp.y)

	for _, tree := range g.gameMap.trees {
		tree.Draw(screen)
	}
	for _, treasure := range g.gameMap.treasures {
		treasure.Draw(screen)
	}
	g.obelisk.Draw(screen)
	g.player.Draw(screen)
}

func (g *GameScene) Exit() {

}
