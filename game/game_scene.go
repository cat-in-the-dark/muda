package sszb

import (
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"sort"
)

type GameScene struct {
	vp            *Viewport
	player        *Player
	mapGenerator  *MapGenerator
	gameMap       *Map
	treasureCount []int
	treasureLeft  []int
	hud           *Hud
	status        *Status
	sm            *lib.SceneManager
}

func NewGameScene(sm *lib.SceneManager) *GameScene {
	vp := NewViewport()
	return &GameScene{
		vp:            vp,
		player:        NewPlayer(vp),
		mapGenerator:  NewMapGenerator(20, vp),
		treasureCount: make([]int, TreasureTypes),
		treasureLeft:  make([]int, TreasureTypes),
		hud:           NewHud(),
		status:        NewStatus(),
		sm:            sm,
	}
}

func (g *GameScene) Activate() {
	g.gameMap = g.mapGenerator.Generate()
	g.treasureLeft = g.gameMap.treasureTotal
	g.status.Show(&g.treasureCount, &g.gameMap.treasureTotal)
	g.hud.showHelp()
}

func (g *GameScene) Update() {
	g.player.Update()
	g.checkCollisions()
	g.hud.Update()
	g.status.Update()
	if g.checkVictory() {
		g.sm.ChangeScene(GameEndName)
	}
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

	g.hud.Draw(screen)
	g.status.Draw(screen)
}

func (g *GameScene) Exit() {

}

func (g *GameScene) checkCollisions() {
	trCollected := false
	trIndex := -1
	trType := -1
	obOffer := false
	var obType int32 = -1
	for i, treasure := range g.gameMap.treasures {
		if Collide(treasure, g.player) {
			trCollected, trIndex, trType = g.collideWithTreasure(treasure, i)
		}
	}
	for i, obelisk := range g.gameMap.obelisks {
		if Collide(obelisk, g.player) {
			obOffer, obType = g.collideWithObelisk(obelisk, i)
		}
	}
	if trCollected && trIndex != -1 {
		g.gameMap.treasures = removeTreasure(g.gameMap.treasures, trIndex)
		g.treasureCount[trType] = g.treasureCount[trType] + 1
		g.treasureLeft[trType] = g.treasureLeft[trType] - 1

		log.Printf("Collected treasures: %v", g.treasureCount)
		log.Printf("Treasures left: %v", g.treasureLeft)
	}
	if obOffer {
		log.Printf("Offered %d treasures to obelisk of type %d", g.treasureCount[obType], obType)
		g.treasureCount[obType] = 0

		for _, obelisk := range g.gameMap.obelisks {
			if obelisk.treasureType == obType {
				obelisk.texture = ObeliskDoneTexture
				obelisk.treasureType = FinalType
			}
		}

		log.Printf("Collected treasures: %v", g.treasureCount)
		log.Printf("Treasures left: %v", g.treasureLeft)
	}
}

func (g *GameScene) collideWithTreasure(treasure *Treasure, index int) (bool, int, int) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.player.treasures = append(g.player.treasures, treasure)
		log.Printf("Added treasureId %d to player inventory", treasure.id)
		return true, index, int(treasure.treasureType)
	}

	g.hud.ShowReset(NewMessage(SeeTreasure[treasure.treasureType], 1))

	return false, -1, -1
}

func (g *GameScene) collideWithObelisk(obelisk *Obelisk, index int) (bool, int32) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		return true, obelisk.treasureType
	}

	g.hud.ShowReset(NewMessage(SeeObelisk[obelisk.treasureType], 1))

	return false, -1
}

func (g *GameScene) checkVictory() bool {
	victory := true
	for i := 0; i < TreasureTypes; i++ {
		collected := g.treasureCount[i]
		remaining := g.treasureLeft[i]
		victory = victory && collected == remaining
	}
	return victory
}

func removeTreasure(t []*Treasure, i int) []*Treasure {
	t[len(t)-1], t[i] = t[i], t[len(t)-1]
	return t[:len(t)-1]
}
