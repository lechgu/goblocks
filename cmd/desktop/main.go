package main

import (
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/lechgu/goblocks/internal/blocks"
)

func main() {
	b := blocks.New()
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Blocks")
	if err := ebiten.RunGame(b); err != nil {
		log.Fatal(err)
	}
}
