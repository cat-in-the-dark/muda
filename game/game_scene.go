package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type GameScene struct {
	vp           *Viewport
	player       *Player
	mapGenerator *MapGenerator
	gameMap      *Map
}

func NewGameScene() *GameScene {
	vp := NewViewport()
	return &GameScene{
		vp:           vp,
		player:       NewPlayer(vp),
		mapGenerator: NewMapGenerator(20, vp),
	}
}

func (g *GameScene) Activate() {
	g.gameMap = g.mapGenerator.Generate()
}

func (g *GameScene) Update() {
	g.player.Update()
	g.checkCollisions()
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(ColorBack)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-g.vp.x, -g.vp.y)

	for _, tree := range g.gameMap.trees {
		tree.Draw(screen)
	}
	for _, obelisk := range g.gameMap.obelisks {
		obelisk.Draw(screen)
	}
	for _, treasure := range g.gameMap.treasures {
		treasure.Draw(screen)
	}
	g.player.Draw(screen)
}

func (g *GameScene) Exit() {

}

func (g *GameScene) checkCollisions() {
	for _, treasure := range g.gameMap.treasures {
		if Collide(treasure, g.player) {
			g.collideWithTreasure(treasure)
		}
	}
	for _, obelisk := range g.gameMap.obelisks {
		if Collide(obelisk, g.player) {
			g.collideWithObelisk(obelisk)
		}
	}
}

func (g *GameScene) collideWithTreasure(treasure *Treasure) {
	log.Printf("Collide %v", treasure)
}

func (g *GameScene) collideWithObelisk(obelisk *Obelisk) {
	log.Printf("Collide %v", obelisk)
}
