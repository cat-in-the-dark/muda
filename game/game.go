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
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, Tutorial1SceneName, sceneManager)).
		With(lib.NewDelayScene(2, Tutorial1SceneName, sceneManager))

	t1 := lib.NewComboScene().
		With(lib.NewTextureScene(Tutorial1Texture, nil)).
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, Tutorial2SceneName, sceneManager))
	t2 := lib.NewComboScene().
		With(lib.NewTextureScene(Tutorial2Texture, nil)).
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, Tutorial3SceneName, sceneManager))
	t3 := lib.NewComboScene().
		With(lib.NewTextureScene(Tutorial3Texture, nil)).
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, NightSceneName, sceneManager))
	dayScene := lib.NewComboScene().
		With(NewDayScene()).
		With(lib.NewKeyAwaitScene(ebiten.KeyEscape, LogoSceneName, sceneManager))
	nightScene := lib.NewComboScene().
		With(NewNightScene()).
		With(lib.NewKeyAwaitScene(ebiten.KeyEscape, LogoSceneName, sceneManager))
	gameWinScene := lib.NewComboScene().
		With(lib.NewTextureScene(GameWinTexture, nil)).
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, LogoSceneName, sceneManager)).
		With(lib.NewKeyAwaitScene(ebiten.KeyEscape, LogoSceneName, sceneManager))
	gameOverScene := lib.NewComboScene().
		With(lib.NewTextureScene(GameOverTexture, nil)).
		With(lib.NewKeyAwaitScene(ebiten.KeySpace, LogoSceneName, sceneManager)).
		With(lib.NewKeyAwaitScene(ebiten.KeyEscape, LogoSceneName, sceneManager))

	sceneManager.Register(LogoSceneName, logo)
	sceneManager.Register(Tutorial1SceneName, t1)
	sceneManager.Register(Tutorial2SceneName, t2)
	sceneManager.Register(Tutorial3SceneName, t3)
	sceneManager.Register(DaySceneName, dayScene)
	sceneManager.Register(NightSceneName, nightScene)
	sceneManager.Register(GameOverSceneName, gameOverScene)
	sceneManager.Register(GameWinSceneName, gameWinScene)

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
