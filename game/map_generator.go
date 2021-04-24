package sszb

import (
	"log"
	"math"
	"math/rand"
)

const (
	CellRows     = 3
	CellColumns  = 3
	TreesPerCell = 20
	CellWidth    = 340
	CellHeight   = 192
	CellMargin   = 16
	TreeDistance = 32
)

var cells [CellRows * CellColumns]*Cell

type Map struct {
	trees []*Tree
}

type MapGenerator struct {
	seed int64
	vp   *Viewport
}

func NewMapGenerator(seed int64, vp *Viewport) *MapGenerator {
	return &MapGenerator{seed, vp}
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
		trees[i] = NewTree(pos, mg.vp)
	}
	return trees
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

	return &Map{
		trees: trees,
	}
}
