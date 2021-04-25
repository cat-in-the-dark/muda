package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type VictoryScene struct{}

func NewVictoryScene() *VictoryScene {
	return &VictoryScene{}
}

func (vs *VictoryScene) Activate() {
}

func (vs *VictoryScene) Update() {
}

func (vs *VictoryScene) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.54, 0.54)
	screen.DrawImage(VictoryScreenTexture, opt)
	text.Draw(screen, "Нажмите Esc, чтобы вернуться на главный экран", DefaultFont, 215, 535, ColorText)
}

func (vs *VictoryScene) Exit() {
}
