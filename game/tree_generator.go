package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

const (
	CellRows     = 3
	CellColumns  = 3
	TreesPerCell = 10
	CellWidth    = 340
	CellHeight   = 192
)

var cells [CellRows * CellColumns]*Cell

type Cell struct {
	x      int32
	y      int32
	width  int32
	height int32
	trees  []*Tree
}

type Tree struct {
	texture  *ebiten.Image
	position *Vector2
}

func NewCell(x, y, width, height int32) *Cell {
	return &Cell{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func NewTree(position *Vector2) *Tree {
	return &Tree{
		texture:  TreeTexture,
		position: position,
	}
}

type TreeGenerator struct {
	seed int64
}

func NewTreeGenerator(seed int64) *TreeGenerator {
	return &TreeGenerator{seed}
}

func GenerateTrees(size int, cell *Cell) []*Tree {
	trees := make([]*Tree, size)
	for i := range trees {
		x := GenerateCoordinate(cell.x, cell.x+cell.width)
		y := GenerateCoordinate(cell.y, cell.y+cell.height)
		trees[i] = NewTree(NewVector2(x, y))
	}
	return trees
}

func GenerateCoordinate(min, max int32) float64 {
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}

func (tg *TreeGenerator) Activate() {
	rand.Seed(tg.seed)
	for i := 0; i < CellRows; i++ {
		for j := 0; j < CellColumns; j++ {
			cell := NewCell(int32(j * CellWidth), int32(i * CellHeight), int32(CellWidth), int32(CellHeight))
			cell.trees = GenerateTrees(TreesPerCell, cell)
			cells[j+i*CellColumns] = cell
		}
	}
}

func (tg *TreeGenerator) Update() {

}

func (tg *TreeGenerator) Draw(screen *ebiten.Image) {
	trees := make([]*Tree, 0)
	for _, cell := range cells {
		for _, tree := range cell.trees {
			trees = append(trees, tree)
		}
	}
	for _, tree := range trees {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(tree.position.x, tree.position.y)
		screen.DrawImage(tree.texture, opt)
	}
}

func (tg *TreeGenerator) Exit() {

}
