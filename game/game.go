package sszb

import (
	"github.com/cat-in-the-dark/muda/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *lib.SceneManager
}

func NewGame() (*Game, error) {
	sceneManager := lib.NewSceneManager()

	logo := lib.NewComboScene().
		With(lib.NewTextureScene(LogoTexture, nil)).
		With(lib.NewDelayScene(2, LogoSceneName, sceneManager))

	sceneManager.Register(LogoSceneName, logo)

	sceneManager.ChangeScene(LogoSceneName)

	g := &Game{
		sceneManager: sceneManager,
	}
	return g, nil
}

func (g Game) Update() error {
	g.sceneManager.Update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
