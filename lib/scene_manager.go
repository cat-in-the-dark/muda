package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Scene interface {
	Activate()
	Update()
	Draw(screen *ebiten.Image)
	Exit()
}

type SceneManager struct {
	scenes       map[string]Scene
	currentScene string
	nextScene    string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes:       make(map[string]Scene),
		currentScene: "",
		nextScene:    "",
	}
}

// ChangeScene schedules a scene swap on the next Update.
func (sc *SceneManager) ChangeScene(name string) {
	log.Printf("Scheduled transition from '%s' to '%s'", sc.currentScene, name)
	sc.nextScene = name
}

func (sc *SceneManager) Update() {
	if sc.currentScene != sc.nextScene {
		sc.safeCall(sc.currentScene, Scene.Exit)
		sc.safeCall(sc.nextScene, Scene.Activate)
		sc.currentScene = sc.nextScene
	}
	sc.safeCall(sc.currentScene, Scene.Update)
}

func (sc *SceneManager) Draw(screen *ebiten.Image) {
	sc.safeCall(sc.currentScene, func(scene Scene) {
		scene.Draw(screen)
	})
}

func (sc *SceneManager) Register(name string, scene Scene) {
	sc.scenes[name] = scene
}

func (sc *SceneManager) safeCall(name string, action func(Scene)) {
	scene, ok := sc.scenes[name]
	if !ok {
		log.Printf("Scene '%s' doesn't exist", name)
		return
	}
	action(scene)
}
