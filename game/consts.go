package sszb

import (
	"github.com/cat-in-the-dark/muda/assets"
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenHeight       = 768
	ScreenWidth        = 1366
	LogoSceneName      = "LOGO"
	Tutorial1SceneName = "TUTORIAL_1"
	Tutorial2SceneName = "TUTORIAL_2"
	Tutorial3SceneName = "TUTORIAL_3"
	GameScene          = "GAME_SCENE"
	GameEnd            = "GAME_END"
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
