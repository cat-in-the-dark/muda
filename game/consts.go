package sszb

import (
	"encoding/hex"
	"github.com/cat-in-the-dark/muda/assets"
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
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
	ColorBack    = HexToColor("f0f6f0")
	ColorHudLine = HexToColor("000000")
	ColorHudBody = HexToColor("f0f6f0")

	LogoTexture    *ebiten.Image
	ObeliskTexture *ebiten.Image

	PlayerIdleTexture  *ebiten.Image
	PlayerDownTexture  *ebiten.Image
	PlayerUpTexture    *ebiten.Image
	PlayerRightTexture *ebiten.Image

	FaceTexture *ebiten.Image

	TreeTextures     []*ebiten.Image
	TreasureTextures []*ebiten.Image

	Animator        *lib.AnimationSystem
	PlayerIdleAnim  *lib.Animation
	PlayerDownAnim  *lib.Animation
	PlayerUpAnim    *lib.Animation
	PlayerRightAnim *lib.Animation
)

// LoadAssets loads textures and other global variables into the scope
// Must be called before NewGame
func LoadAssets() {
	Animator = lib.NewAnimationSystem()
	LogoTexture = assets.LoadImage("logo.png")

	TreasureTextures = []*ebiten.Image{
		assets.LoadImage("crane.png"),
		assets.LoadImage("frog.png"),
		assets.LoadImage("squirrel.png"),
		assets.LoadImage("star.png"),
		assets.LoadImage("elephant.png"),
		assets.LoadImage("dragon.png"),
	}

	FaceTexture = assets.LoadImage("face.png")

	TreeTextures = []*ebiten.Image{
		assets.LoadImage("tree1.png"),
		assets.LoadImage("tree2.png"),
	}

	PlayerIdleTexture = assets.LoadImage("stay.png")
	PlayerDownTexture = assets.LoadImage("walking_down.png")
	PlayerUpTexture = assets.LoadImage("walking_up.png")
	PlayerRightTexture = assets.LoadImage("walking_right.png")

	ObeliskTexture = assets.LoadImage("obelisk.png")

	playerIdleSheet := lib.NewSpriteSheet(PlayerIdleTexture, 64, 64)
	playerDownSheet := lib.NewSpriteSheet(PlayerDownTexture, 64, 64)
	playerUpSheet := lib.NewSpriteSheet(PlayerUpTexture, 64, 64)
	playerRightSheet := lib.NewSpriteSheet(PlayerRightTexture, 64, 64)

	PlayerIdleAnim = Animator.NewLooping(playerIdleSheet, 36, []int{0, 1})
	PlayerDownAnim = Animator.NewLooping(playerDownSheet, 8, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	PlayerUpAnim = Animator.NewLooping(playerUpSheet, 8, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	PlayerRightAnim = Animator.NewLooping(playerRightSheet, 8, []int{0, 1, 2, 3, 4, 5, 6})
}

func HexToColor(hexColor string) color.Color {
	b, err := hex.DecodeString(hexColor)
	if err != nil {
		log.Fatal(err)
	}

	return color.RGBA{b[0], b[1], b[2], 255}
}
