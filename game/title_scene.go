package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type TitleScene struct{}

func NewTitleScene() *TitleScene {
	return &TitleScene{}
}

func (ts *TitleScene) Activate() {
}

func (ts *TitleScene) Update() {
}

func (ts *TitleScene) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.54, 0.54)
	screen.DrawImage(TitleScreenTexture, opt)
	text.Draw(screen, "Нажмите Enter\nчтобы начать игру", DefaultFont, 690, 220, ColorText)
}

func (ts *TitleScene) Exit() {
}

