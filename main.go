package main

import (
	_ "image/png"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
