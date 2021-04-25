package main

import (
	"log"

	game "github.com/cat-in-the-dark/muda/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game.LoadAssets()
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	//ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Manager's Ubelievable Daydream Adventures")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
