package sszb

import (
	"log"
	"math"
	"math/rand"
)

const (
	TreasureTypes = 6
	FinalType     = TreasureTypes // alias
	CellRows      = 3
	CellColumns   = 3
	TreesPerCell  = 35
	TreasureCount = 30
	ObeliskCount  = 12
	CellWidth     = 1600
	CellHeight    = 900
	CellMargin    = 32
	TreeDistance  = 100
	MapWidth      = CellColumns * CellWidth
	MapHeight     = CellRows * CellHeight
)

var cells [CellRows * CellColumns]*Cell

type Map struct {
	trees         []*Tree
	treasures     []*Treasure
	obelisks      []*Obelisk
	treasureTotal []int
}

type MapGenerator struct {
	seed          int64
	vp            *Viewport
	treasureTotal []int
}

func NewMapGenerator(seed int64, vp *Viewport) *MapGenerator {
	return &MapGenerator{seed, vp, make([]int, TreasureTypes)}
}

func (mg *MapGenerator) GenerateTrees(size int, cell *Cell) []*Tree {
	trees := make([]*Tree, size)
	for i := range trees {
		treeGenerated := false
		var pos *Vector2 = nil
		for {
			x := GenerateCoordinate(cell.x, cell.x+cell.width)
			y := GenerateCoordinate(cell.y, cell.y+cell.height)
			if CheckTreePosition(x, y, trees) {
				pos = NewVector2(x, y)
				treeGenerated = true
			}
			if treeGenerated && pos != nil {
				break
			}
		}

		log.Printf("Generated tree at x:%f, y:%f", pos.x, pos.y)
		trees[i] = NewTree(pos, mg.vp, rand.Int31n(int32(len(TreeTextures))))
	}
	return trees
}

func (mg *MapGenerator) GenerateTreasures(size int) []*Treasure {
	treasures := make([]*Treasure, size)
	for i := range treasures {
		x := GenerateCoordinate(CellMargin, MapWidth-CellMargin)
		y := GenerateCoordinate(CellMargin, MapHeight-CellMargin)
		pos := NewVector2(x, y)
		log.Printf("Generated treasure at x:%f, y:%f", pos.x, pos.y)
		typ := rand.Int31n(TreasureTypes)
		treasures[i] = NewTreasure(pos, mg.vp, typ)
		mg.treasureTotal[typ] = mg.treasureTotal[typ] + 1
	}
	return treasures
}

func (mg *MapGenerator) GenerateObelisks(size int) []*Obelisk {
	obelisks := make([]*Obelisk, size)
	for i := range obelisks {
		x := GenerateCoordinate(CellMargin, MapWidth-CellMargin)
		y := GenerateCoordinate(CellMargin, MapHeight-CellMargin)
		pos := NewVector2(x, y)
		log.Printf("Generated obelisk at x:%f, y:%f", pos.x, pos.y)
		obelisks[i] = NewObelisk(pos, mg.vp, rand.Int31n(int32(TreasureTypes)))
	}
	return obelisks
}

func CheckTreePosition(x, y float64, trees []*Tree) bool {
	noOverlap := true
	for i := 0; i < cap(trees); i++ {
		tree := trees[i]
		if tree == nil {
			break
		}
		pos := tree.pos
		dist := math.Sqrt(math.Pow(x-pos.x, 2) + math.Pow(y-pos.y, 2))
		noOverlap = noOverlap && dist > TreeDistance
		if !noOverlap {
			break
		}
	}
	return noOverlap
}

func GenerateCoordinate(min, max int32) float64 {
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}

func (mg *MapGenerator) Generate() *Map {
	rand.Seed(mg.seed)
	for i := 0; i < CellRows; i++ {
		for j := 0; j < CellColumns; j++ {
			cell := NewCell(
				int32(j*(CellWidth+CellMargin)),
				int32(i*(CellHeight+CellMargin)),
				int32(CellWidth-CellMargin),
				int32(CellHeight-CellMargin))
			cell.trees = mg.GenerateTrees(TreesPerCell, cell)
			cells[j+i*CellColumns] = cell
		}
	}
	trees := make([]*Tree, 0)
	for _, cell := range cells {
		for _, tree := range cell.trees {
			trees = append(trees, tree)
		}
	}
	obelisks := mg.GenerateObelisks(ObeliskCount)
	treasures := mg.GenerateTreasures(TreasureCount)

	return &Map{
		trees:         trees,
		treasures:     treasures,
		obelisks:      obelisks,
		treasureTotal: mg.treasureTotal,
	}
}
