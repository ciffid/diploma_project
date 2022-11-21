package main

import (
	"dota3/assets"
	"dota3/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	assets.Init()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
