package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ComboScene is a Scene which is a composition of functions
type ComboScene struct {
	scenes []Scene
}

func NewComboScene() *ComboScene {
	return &ComboScene{
		scenes: make([]Scene, 0),
	}
}

func (sc *ComboScene) With(scene Scene) *ComboScene {
	sc.scenes = append(sc.scenes, scene)
	return sc
}

func (sc *ComboScene) Activate() {
	for _, scene := range sc.scenes {
		scene.Activate()
	}
}

func (sc *ComboScene) Update() {
	for _, scene := range sc.scenes {
		scene.Update()
	}
}

func (sc *ComboScene) Draw(screen *ebiten.Image) {
	for _, scene := range sc.scenes {
		scene.Draw(screen)
	}
}

func (sc *ComboScene) Exit() {
	for _, scene := range sc.scenes {
		scene.Exit()
	}
}

type TextureScene struct {
	texture *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewTextureScene(texture *ebiten.Image, options *ebiten.DrawImageOptions) *TextureScene {
	return &TextureScene{
		texture: texture,
		options: options,
	}
}

func (sc *TextureScene) Activate() {

}

func (sc *TextureScene) Update() {

}

func (sc *TextureScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(sc.texture, sc.options)
}

func (sc *TextureScene) Exit() {

}

type KeyAwaitScene struct {
	key       ebiten.Key
	nextScene string
	manager   *SceneManager
}

func NewKeyAwaitScene(
	key ebiten.Key,
	nextScene string,
	manager *SceneManager,
) *KeyAwaitScene {
	return &KeyAwaitScene{
		key:       key,
		nextScene: nextScene,
		manager:   manager,
	}
}

func (sc *KeyAwaitScene) Activate() {

}

func (sc *KeyAwaitScene) Update() {
	if inpututil.IsKeyJustPressed(sc.key) {
		sc.manager.ChangeScene(sc.nextScene)
	}
}

func (sc *KeyAwaitScene) Draw(screen *ebiten.Image) {

}

func (sc *KeyAwaitScene) Exit() {

}

type DelayScene struct {
	currentTime float64
	delayTime   float64
	nextScene   string
	manager     *SceneManager
}

func NewDelayScene(
	delayTime float64,
	nextScene string,
	manager *SceneManager,
) *DelayScene {
	return &DelayScene{
		currentTime: 0,
		delayTime:   delayTime,
		nextScene:   nextScene,
		manager:     manager,
	}
}

func (sc* DelayScene) Activate() {
	sc.currentTime = 0
}

func (sc* DelayScene) Update() {
	if ebiten.CurrentTPS() < 1 {
		// protect agains zero division
		return
	}
	sc.currentTime += 1 / ebiten.CurrentTPS()

	if sc.currentTime >= sc.delayTime {
		sc.manager.ChangeScene(sc.nextScene)
	}
}

func (sc* DelayScene) Draw(screen *ebiten.Image) {

}

func (sc* DelayScene) Exit() {

}

type DrawScene struct {
	draw func(screen *ebiten.Image)
}

func NewDrawScene(draw func(screen *ebiten.Image)) *DrawScene {
	return &DrawScene{
		draw: draw,
	}
}

func (sc* DrawScene) Activate() {
}

func (sc* DrawScene) Update() {
}

func (sc* DrawScene) Draw(screen *ebiten.Image) {
	sc.draw(screen)
}

func (sc* DrawScene) Exit() {
}



