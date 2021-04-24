package sszb

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type NightScene struct {

}

func NewNightScene() *NightScene {
	return &NightScene{}
}

func (sc *NightScene) Activate() {

}

func (sc *NightScene) Update() {
	Animator.Update()
}

func (sc *NightScene) Draw(screen *ebiten.Image) {
	drawNightBackground(screen)

	screen.DrawImage(ChikaWalkAnim.GetFrame(), nil)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 0)
	screen.DrawImage(ChikaAttackAnim.GetFrame(), op)
}

func (sc *NightScene) Exit() {

}

func drawNightBackground(screen *ebiten.Image) {
	screen.DrawImage(NightBgTexture, nil)
}
