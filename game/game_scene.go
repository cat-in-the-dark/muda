package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"sort"
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

	drawables := make([]Drawable, len(g.gameMap.trees)+len(g.gameMap.obelisks)+len(g.gameMap.treasures)+1)

	i := 0
	for _, tree := range g.gameMap.trees {
		drawables[i] = tree
		i++
	}
	for _, obelisk := range g.gameMap.obelisks {
		drawables[i] = obelisk
		i++
	}
	for _, treasure := range g.gameMap.treasures {
		drawables[i] = treasure
		i++
	}
	drawables[i] = g.player

	sort.Slice(drawables, func(i, j int) bool {
		return drawables[i].GetPos().y < drawables[j].GetPos().y
	})
	for _, drawable := range drawables {
		drawable.Draw(screen)
	}
}

func (g *GameScene) Exit() {

}

func (g *GameScene) checkCollisions() {
	trCollected := false
	trIndex := -1
	for i, treasure := range g.gameMap.treasures {
		if Collide(treasure, g.player) {
			trCollected, trIndex = g.collideWithTreasure(treasure, i)
		}
	}
	for i, obelisk := range g.gameMap.obelisks {
		if Collide(obelisk, g.player) {
			g.collideWithObelisk(obelisk, i)
		}
	}
	if trCollected && trIndex != -1 {
		g.gameMap.treasures = removeTreasure(g.gameMap.treasures, trIndex)
	}
}

func (g *GameScene) collideWithTreasure(treasure *Treasure, index int) (bool, int) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.player.treasures = append(g.player.treasures, treasure)
		log.Printf("Added treasureId %d to player inventory", treasure.id)
		return true, index
	}
	return false, -1
}

func (g *GameScene) collideWithObelisk(obelisk *Obelisk, index int) {
	//log.Printf("Collide %v", obelisk)
}

func removeTreasure(t []*Treasure, i int) []*Treasure {
	t[len(t)-1], t[i] = t[i], t[len(t)-1]
	return t[:len(t)-1]
}
