package sszb

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Status struct {
	pos       *Vector2
	size      *Vector2
	trCount   *[]int
	trTotal   *[]int
	lw        float64 // line width
	txtOffset *Vector2
}

func NewStatus() *Status {
	return &Status{
		pos:       NewVector2(0, 0),
		size:      NewVector2(1024, 48),
		lw:        2,
		txtOffset: NewVector2(8, 10+FontSize),
	}
}

func (st *Status) Show(trCount *[]int, trTotal *[]int) {
	st.trCount = trCount
	st.trTotal = trTotal
}

func (st *Status) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, st.pos.x, st.pos.y, st.size.x, st.size.y, ColorHudLine)
	ebitenutil.DrawRect(screen, st.pos.x+st.lw, st.pos.y+st.lw, st.size.x-st.lw*2, st.size.y-st.lw*2, ColorHudBody)

	for i, count := range *st.trCount {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(st.pos.x+st.txtOffset.x*float64(i)*15+166, st.pos.y+8)
		screen.DrawImage(TreasureTextures[i], opt)
		text.Draw(
			screen,
			fmt.Sprintf("%2d/%2d", count, (*st.trTotal)[i]),
			DefaultFont,
			int(st.pos.x+(st.txtOffset.x*float64(i)*15)+206),
			int(st.pos.y+st.txtOffset.y),
			ColorText)
	}
}

func (st *Status) Update() {
}
