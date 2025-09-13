package main

import (
	"fmt"
	"log"
	"math"
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
					//fmt.Printf("%c\n", selectedPiece.nameCode)
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
		//fmt.Printf("Selected square at: %d, %d", selectedSquare.x, selectedSquare.y)

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
	case 'K':
		if canKnightMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	case 'B':
		if canBishopMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	case 'Q':
		if canQueenMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	case 'X':
		if canKingMove(p, selectedSquare) {
			movePiece(p, selectedSquare)
		}
	}
}

func canPawnMove(p *Piece, selectedSquare *Square) bool {
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

	moveHorizontal := false
	if p.posY == selectedSquare.y {
		moveHorizontal = true
	}
	fmt.Printf("Selected piece: %d, %d, selected square: %d, %d\n", selectedPiece.posX, selectedPiece.posY, selectedSquare.x, selectedSquare.y)
	//TODO: Change square select method, should piece be selected first then square
	if (selectedSquare.x == p.posX && selectedSquare.y == p.posY) || (selectedSquare.x != p.posX && selectedSquare.y != p.posY) {
		return false
	}

	if moveHorizontal && selectedSquare.x > p.posX && selectedSquare.y == p.posY { // Right Move
		for j := p.posX + 1; j <= selectedSquare.x; j++ {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if moveHorizontal && selectedSquare.x < p.posX && selectedSquare.y == p.posY { // Left Move
		for j := p.posX - 1; j >= selectedSquare.x; j-- {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y < p.posY && selectedSquare.x == p.posX { // Up Move
		for j := p.posY - 1; j >= selectedSquare.y; j-- {
			if board[j][p.posX] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y > p.posY && selectedSquare.x == p.posX { // Down Move
		for j := p.posY + 1; j <= selectedSquare.y; j++ {
			if board[j][p.posX] != nil {
				return false
			}
		}
	}
	return true
}

// This function can be implemented with a bunch of if statements without using extra space or loop, but this is much clearer I think
func canKnightMove(p *Piece, selectedSquare *Square) bool {
	type pair struct {
		dx int
		dy int
	}
	var possibleMoves = [8]*pair{
		{1, -2}, {-1, -2}, // Upward move
		{1, 2}, {-1, 2}, // Downward move
		{2, -1}, {2, 1}, // Right move
		{-2, -1}, {-2, 1}, // Left move
	}
	//TODO: Check if there is a piece at selected square
	for _, move := range possibleMoves {
		if p.posX+move.dx == selectedSquare.x && p.posY+move.dy == selectedSquare.y {
			return true
		}
	}
	return false
}

func canBishopMove(p *Piece, selectedSquare *Square) bool {
	b := false
	dx, dy := 0, 0

	// Diagonal search for selected square
	for i := 1; i <= 6; i++ {
		if p.posX+i == selectedSquare.x && p.posY+i == selectedSquare.y {
			b = true
			dx = i
			dy = i
		} else if p.posX+i == selectedSquare.x && p.posY-i == selectedSquare.y {
			b = true
			dx = i
			dy = -i
		} else if p.posX-i == selectedSquare.x && p.posY+i == selectedSquare.y {
			b = true
			dx = -i
			dy = i
		} else if p.posX-i == selectedSquare.x && p.posY-i == selectedSquare.y {
			b = true
			dx = -i
			dy = -i
		}
	}

	if !b {
		return false
	}

	if dx > 0 && dy > 0 {
		for i := 1; i <= dx; i++ {
			if board[p.posY+i][p.posX+i] != nil {
				return false
			}
		}
	} else if dx > 0 && dy < 0 {
		for i := 1; i <= dx; i++ {
			if board[p.posY-i][p.posX+i] != nil {
				return false
			}
		}
	} else if dx < 0 && dy > 0 {
		for i := 1; i <= dy; i++ {
			if board[p.posY+i][p.posX-i] != nil {
				return false
			}
		}
	} else {
		for i := -1; i >= dx; i-- {
			if board[p.posY+i][p.posX+i] != nil {
				return false
			}
		}
	}
	return true
}

func canQueenMove(p *Piece, selectedSquare *Square) bool {
	if canRookMove(p, selectedSquare) || canBishopMove(p, selectedSquare) {
		return true
	}
	return false
}

func canKingMove(p *Piece, selectedSquare *Square) bool {
	if p.posY == selectedSquare.y && (p.posX+1 == selectedSquare.x || p.posX-1 == selectedSquare.x) {
		return true
	} else if p.posX == selectedSquare.x && (p.posY+1 == selectedSquare.y || p.posY-1 == selectedSquare.y) {
		return true
	} else if math.Abs(float64(p.posX-selectedSquare.x)) == 1 && math.Abs(float64(p.posY-selectedSquare.y)) == 1 {
		return true
	}
	return false
}

func movePiece(p *Piece, selectedSquare *Square) {
	prevX, prevY := p.posX, p.posY // Initial positions for clearing the board
	p.posX = selectedSquare.x
	p.posY = selectedSquare.y

	board[p.posY][p.posX] = p
	board[prevY][prevX] = nil

	for _, line := range board {
		for _, p := range line {
			if p == nil {
				fmt.Printf("- ")
				continue
			}
			fmt.Printf("%c, ", p.nameCode)
		}
		fmt.Println("")
	}
	fmt.Println("")

	p.isMoved = true
	selectedPiece = nil
	selectedSquare = nil
}
