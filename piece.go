package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Piece struct {
	posX, posY int
	img        *ebiten.Image
	nameCode   uint8
	color      bool // false -> White, true -> Black
	isMoved    bool
}

type Square struct {
	x int
	y int
}

var (
	selectedPiece  *Piece
	selectedSquare Square
)

// init initializes images for the pieces
func init() {
	cd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range board {
		for _, piece := range line {
			if piece == nil {
				continue
			}
			if piece.nameCode == 'R' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_04.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'K' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_03.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'B' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_02.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'Q' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_01.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'X' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_00.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'P' && !piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_05.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'R' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_10.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'K' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_09.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'B' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_08.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'Q' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_07.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'X' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_06.png"))
				if err != nil {
					log.Fatal(err)
				}
			} else if piece.nameCode == 'P' && piece.color {
				piece.img, _, err = ebitenutil.NewImageFromFile(filepath.Join(cd, "/ChessPieces/piece_11.png"))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

// DrawPieces draws the pieces into center of their current square
func DrawPieces(screen *ebiten.Image) {
	for _, line := range board {
		for _, piece := range line {
			if piece != nil && piece.img != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.5, 0.5)

				// Image size
				w, h := piece.img.Size()
				tw := float64(w) * 0.5
				th := float64(h) * 0.5

				//Finding current square position
				x := float64(piece.posX * squareSize)
				y := float64(piece.posY * squareSize)

				//Centering the square
				x += (float64(squareSize) - tw) / 2
				y += (float64(squareSize) - th) / 2

				op.GeoM.Translate(x, y)
				screen.DrawImage(piece.img, op)
			}
		}
	}
}

// CheckSelectedPiece checks if user selected a piece to move
func CheckSelectedPiece() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()

		if x > width {
			return
		}
		x /= squareSize
		y /= squareSize

		for _, line := range board {
			for _, piece := range line {
				if piece != nil && x == piece.posX && y == piece.posY && piece.color == turn {
					selectedPiece = piece
					//fmt.Printf("%c\n", selectedPiece.nameCode)
					return
				}
			}
		}
	}
}
