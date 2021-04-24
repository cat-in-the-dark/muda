package sszb

import (
	"github.com/cat-in-the-dark/muda/assets"
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenHeight  = 576
	ScreenWidth   = 1024
	LogoSceneName = "LOGO"
	GameSceneName = "GAME_SCENE"
	GameEndName   = "GAME_END"
)

var (
	LogoTexture *ebiten.Image

	Animator *lib.AnimationSystem
)

// LoadAssets loads textures and other global variables into the scope
// Must be called before NewGame
func LoadAssets() {
	Animator = lib.NewAnimationSystem()
	LogoTexture = assets.LoadImage("logo.png")
}
