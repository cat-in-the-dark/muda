package sszb

import (
	"encoding/hex"
	"github.com/cat-in-the-dark/muda/assets"
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

const (
	FontSize               = 24
	ScreenHeight           = 576
	ScreenWidth            = 1024
	LogoSceneName          = "LOGO"
	TreeGeneratorSceneName = "TREE_GENERATOR"
	GameSceneName          = "GAME_SCENE"
	GameEndName            = "GAME_END"

	PlayerStartPosX = 2400
	PlayerStartPosY = 1350

	HudWidth   = 512
	HudHeight  = 128
	HudOffsetY = 32
)

var (
	ColorBack    = HexToColor("f0f6f0")
	ColorHudLine = HexToColor("000000")
	ColorHudBody = HexToColor("f0f6f0")
	ColorText    = HexToColor("000000")

	LogoTexture *ebiten.Image

	PlayerIdleTexture  *ebiten.Image
	PlayerDownTexture  *ebiten.Image
	PlayerUpTexture    *ebiten.Image
	PlayerRightTexture *ebiten.Image

	FaceTexture *ebiten.Image

	ObeliskDoneTexture *ebiten.Image

	TreeTextures     []*ebiten.Image
	ObeliskTextures  []*ebiten.Image
	TreasureTextures []*ebiten.Image

	Animator        *lib.AnimationSystem
	PlayerIdleAnim  *lib.Animation
	PlayerDownAnim  *lib.Animation
	PlayerUpAnim    *lib.Animation
	PlayerRightAnim *lib.Animation

	DefaultFont font.Face
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
		assets.LoadImage("bush.png"),
	}

	PlayerIdleTexture = assets.LoadImage("stay.png")
	PlayerDownTexture = assets.LoadImage("walking_down.png")
	PlayerUpTexture = assets.LoadImage("walking_up.png")
	PlayerRightTexture = assets.LoadImage("walking_right.png")

	ObeliskDoneTexture = assets.LoadImage("obelisk.png")
	ObeliskTextures = []*ebiten.Image{
		assets.LoadImage("obelisk_crane.png"),
		assets.LoadImage("obelisk_frog.png"),
		assets.LoadImage("obelisk_squirrel.png"),
		assets.LoadImage("obelisk_star.png"),
		assets.LoadImage("obelisk_elephant.png"),
		assets.LoadImage("obelisk_dragon.png"),
	}

	playerIdleSheet := lib.NewSpriteSheet(PlayerIdleTexture, 64, 64)
	playerDownSheet := lib.NewSpriteSheet(PlayerDownTexture, 64, 64)
	playerUpSheet := lib.NewSpriteSheet(PlayerUpTexture, 64, 64)
	playerRightSheet := lib.NewSpriteSheet(PlayerRightTexture, 64, 64)

	PlayerIdleAnim = Animator.NewLooping(playerIdleSheet, 36, []int{0, 1})
	PlayerDownAnim = Animator.NewLooping(playerDownSheet, 8, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	PlayerUpAnim = Animator.NewLooping(playerUpSheet, 8, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	PlayerRightAnim = Animator.NewLooping(playerRightSheet, 8, []int{0, 1, 2, 3, 4, 5, 6})

	var err error
	tt := assets.LoadFont("cyrillic_pixel.ttf")
	DefaultFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    FontSize,
		DPI:     76,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func HexToColor(hexColor string) color.Color {
	b, err := hex.DecodeString(hexColor)
	if err != nil {
		log.Fatal(err)
	}

	return color.RGBA{b[0], b[1], b[2], 255}
}
