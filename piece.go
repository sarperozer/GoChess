package main

import (
	"fmt"
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

var (
	selectedPiece  *Piece
	selectedSquare Square
)

func init() {
	cd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	// Initializing images
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

// drawPieces draws the pieces into center of their current square
func drawPieces(screen *ebiten.Image) {
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

// checkSelectedPiece checks if user selected a piece to move
func checkSelectedPiece() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		x /= 100
		y /= 100

		for _, line := range board {
			for _, piece := range line {
				if piece != nil && x == piece.posX && y == piece.posY {
					selectedPiece = piece
					fmt.Printf("%c\n", selectedPiece.nameCode)
					return
				}
			}
		}
	}
}

func checkSelectedSquare() {
	if selectedPiece != nil && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		selectedSquare.x = x / 100
		selectedSquare.y = y / 100

		canPieceMoveSelectedSquare(selectedPiece, &selectedSquare)
	}
}

func canPieceMoveSelectedSquare(p *Piece, selectedSquare *Square) {

	switch p.nameCode {
	case 'P':
		if canPawnMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	case 'R':
		if canRookMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	}

}

func canPawnMove(p *Piece, selectedSquare *Square) bool {
	fmt.Printf("Selected piece at : %d, %d, Selected square at: %d, %d", p.posX, p.posY, selectedSquare.x, selectedSquare.y)
	if p.color == false && selectedSquare.y == p.posY-1 && selectedSquare.x == p.posX {
		return true
	} else if p.color == false && selectedSquare.y == p.posY-2 && selectedSquare.x == p.posX && !p.isMoved {
		return true
	} else if p.color == true && selectedSquare.y == p.posY+1 && selectedSquare.x == p.posX {
		return true
	} else if p.color == true && selectedSquare.y == p.posY+2 && selectedSquare.x == p.posX && !p.isMoved {
		return true
	}
	return false
}

func canRookMove(p *Piece, selectedSquare *Square) bool {
	fmt.Printf("Selected piece at : %d, %d, Selected square at: %d, %d", p.posX, p.posY, selectedSquare.x, selectedSquare.y)
	moveHorizontal := false
	if p.posX == selectedSquare.x {
		moveHorizontal = true
	}
	//TODO: Check if there is friendly piece at selected square
	if moveHorizontal && selectedSquare.x > p.posX { // Right Move
		for j := p.posX + 1; j <= selectedSquare.x; j++ {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if moveHorizontal && selectedSquare.x < p.posX { // Left Move
		for j := p.posX - 1; j >= selectedSquare.x; j-- {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y < p.posX { // Up Move
		for j := p.posY - 1; j >= selectedSquare.y; j-- {
			if board[j][p.posX] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y > p.posY { // Down Move
		for j := p.posY + 1; j <= selectedSquare.y; j++ {
			if board[j][p.posX] != nil {
				return false
			}
		}
	}
	return true
}

func movePiece(p *Piece, selectedSquare *Square) {
	p.posX = selectedSquare.x
	p.posY = selectedSquare.y

	p.isMoved = true
	selectedPiece = nil
	selectedSquare = nil
}
