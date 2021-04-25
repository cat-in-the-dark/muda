package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"strconv"
)

type Status struct {
	pos       *Vector2
	size      *Vector2
	trCount   *[]int
	lw        float64 // line width
	txtOffset *Vector2
}

func NewStatus() *Status {
	return &Status{
		pos:       NewVector2(524, 0),
		size:      NewVector2(500, 48),
		lw:        2,
		txtOffset: NewVector2(8, 10+FontSize),
	}
}

func (st *Status) Show(trCount *[]int) {
	st.trCount = trCount
}

func (st *Status) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, st.pos.x, st.pos.y, st.size.x, st.size.y, ColorHudLine)
	ebitenutil.DrawRect(screen, st.pos.x+st.lw, st.pos.y+st.lw, st.size.x-st.lw*2, st.size.y-st.lw*2, ColorHudBody)

	for i, count := range *st.trCount {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(st.pos.x+st.txtOffset.x*float64(i)*10+16, st.pos.y+8)
		screen.DrawImage(TreasureTextures[i], opt)
		text.Draw(
			screen,
			strconv.Itoa(count),
			DefaultFont,
			int(st.pos.x+(st.txtOffset.x*float64(i)*10)+56),
			int(st.pos.y+st.txtOffset.y),
			ColorText)
	}
	//text.Draw(screen, msg.text, msg.font, int(st.pos.x+st.txtOffset.x), int(st.pos.y+st.txtOffset.y), ColorText)
}

func (st *Status) Update() {
}
