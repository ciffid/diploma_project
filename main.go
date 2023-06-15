package main

import (
	"DP/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(1100, 1100)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
