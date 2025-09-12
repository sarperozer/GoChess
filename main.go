package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	checkSelectedPiece()
	checkSelectedSquare()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBoard(screen)
	drawPieces(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(width, height)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
