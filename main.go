package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const width = 800
const height = 800
const squareSize = 100
const turnBarWidth = 25

type Game struct{}

func (g *Game) Update() error {
	CheckSelectedPiece()
	CheckSelectedSquare()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBoard(screen)
	DrawPieces(screen)
	DrawTurn(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(width+turnBarWidth, height)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func DrawTurn(screen *ebiten.Image) {
	if !turn {
		vector.DrawFilledRect(screen, width, 0, turnBarWidth, height, color.White, false)
	} else {
		vector.DrawFilledRect(screen, width, 0, turnBarWidth, height, color.Black, false)
	}
}
