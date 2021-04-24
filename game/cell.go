package sszb

type Cell struct{
	x      int32
	y      int32
	width  int32
	height int32
	trees  []*Tree
}

func NewCell(x, y, width, height int32) *Cell {
	return &Cell{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}
