package sszb

import (
	"github.com/cat-in-the-dark/muda/assets"
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenHeight           = 576
	ScreenWidth            = 1024
	LogoSceneName          = "LOGO"
	TreeGeneratorSceneName = "TREE_GENERATOR"
	GameSceneName          = "GAME_SCENE"
	GameEndName            = "GAME_END"

	PlayerStartPosX = 128
	PlayerStartPosY = 128
)

var (
	LogoTexture    *ebiten.Image
	PlayerTexture  *ebiten.Image
	BG             *ebiten.Image
	ObeliskTexture *ebiten.Image
	TreeTexture    *ebiten.Image

	Animator *lib.AnimationSystem
)

// LoadAssets loads textures and other global variables into the scope
// Must be called before NewGame
func LoadAssets() {
	Animator = lib.NewAnimationSystem()
	LogoTexture = assets.LoadImage("logo.png")
	PlayerTexture = assets.LoadImage("player.png")
	BG = assets.LoadImage("BG.png")
	TreeTexture = assets.LoadImage("apple_tree2.png")
	ObeliskTexture = assets.LoadImage("obelisk.png")
}
