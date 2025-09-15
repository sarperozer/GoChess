package main

import (
	"fmt"
	"math"
)

var turn = false // false = white move, true = black move

func canPieceMoveSelectedSquare(p *Piece, selectedSquare *Square) {

	switch p.nameCode {
	case 'P':
		if canPawnMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	case 'R':
		if canRookMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	case 'K':
		if canKnightMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	case 'B':
		if canBishopMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	case 'Q':
		if canQueenMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	case 'X':
		if canKingMove(p, selectedSquare) && turn == p.color {
			movePiece(p, selectedSquare)
			turn = !turn
		}
	}
}

func canPawnMove(p *Piece, selectedSquare *Square) bool {
	x, y := selectedSquare.x, selectedSquare.y

	//Checks for if pawn can capture
	if p.color == false && selectedSquare.y == p.posY-1 && selectedSquare.x == p.posX+1 && board[y][x] != nil { // Pawn is white and there is a capture at its up-right
		return true
	} else if p.color == false && selectedSquare.y == p.posY-1 && selectedSquare.x == p.posX-1 && board[y][x] != nil { // Pawn is white and there is a capture at its up-left
		return true
	} else if p.color == true && selectedSquare.y == p.posY+1 && selectedSquare.x == p.posX+1 && board[y][x] != nil { // Pawn is black and there is a capture at its down-right
		return true
	} else if p.color == true && selectedSquare.y == p.posY+1 && selectedSquare.x == p.posX-1 && board[y][x] != nil { // Pawn is black and there is a capture at its down-left
		return true
	}

	//Checks for if pawn can move forward
	if p.color == false && selectedSquare.y == p.posY-1 && selectedSquare.x == p.posX && board[y][x] == nil {
		return true
	} else if p.color == false && selectedSquare.y == p.posY-2 && selectedSquare.x == p.posX && !p.isMoved && board[y][x] == nil {
		return true
	} else if p.color == true && selectedSquare.y == p.posY+1 && selectedSquare.x == p.posX && board[y][x] == nil {
		return true
	} else if p.color == true && selectedSquare.y == p.posY+2 && selectedSquare.x == p.posX && !p.isMoved && board[y][x] == nil {
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

	if (selectedSquare.x == p.posX && selectedSquare.y == p.posY) || (selectedSquare.x != p.posX && selectedSquare.y != p.posY) {
		return false
	}

	if moveHorizontal && selectedSquare.x > p.posX && selectedSquare.y == p.posY { // Right Move
		for j := p.posX + 1; j < selectedSquare.x; j++ {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if moveHorizontal && selectedSquare.x < p.posX && selectedSquare.y == p.posY { // Left Move
		for j := p.posX - 1; j > selectedSquare.x; j-- {
			if board[p.posY][j] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y < p.posY && selectedSquare.x == p.posX { // Up Move
		for j := p.posY - 1; j > selectedSquare.y; j-- {
			if board[j][p.posX] != nil {
				return false
			}
		}
	} else if !moveHorizontal && selectedSquare.y > p.posY && selectedSquare.x == p.posX { // Down Move
		for j := p.posY + 1; j < selectedSquare.y; j++ {
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
		{1, -2}, {-1, -2}, // Upward moves
		{1, 2}, {-1, 2}, // Downward moves
		{2, -1}, {2, 1}, // Right moves
		{-2, -1}, {-2, 1}, // Left moves
	}
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
		if p.posX+i == selectedSquare.x && p.posY+i == selectedSquare.y { // Check for right-upward diagonal move
			b = true
			dx = i
			dy = i
		} else if p.posX+i == selectedSquare.x && p.posY-i == selectedSquare.y { // Check for right-downward diagonal move
			b = true
			dx = i
			dy = -i
		} else if p.posX-i == selectedSquare.x && p.posY+i == selectedSquare.y { // Check for left-upward diagonal move
			b = true
			dx = -i
			dy = i
		} else if p.posX-i == selectedSquare.x && p.posY-i == selectedSquare.y { // Check for left-downward diagonal move
			b = true
			dx = -i
			dy = -i
		}
	}

	if !b {
		return false
	} // Selected square not in the possible moves (result of search) of bishop, return false

	//Checks for if there is a piece on the way to the selected square
	if dx > 0 && dy > 0 {
		for i := 1; i < dx; i++ {
			if board[p.posY+i][p.posX+i] != nil {
				return false
			}
		}
	} else if dx > 0 && dy < 0 {
		for i := 1; i < dx; i++ {
			if board[p.posY-i][p.posX+i] != nil {
				return false
			}
		}
	} else if dx < 0 && dy > 0 {
		for i := 1; i < dy; i++ {
			if board[p.posY+i][p.posX-i] != nil {
				return false
			}
		}
	} else {
		for i := -1; i > dx; i-- {
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
	if p.posY == selectedSquare.y && (p.posX+1 == selectedSquare.x || p.posX-1 == selectedSquare.x) { // Horizontal move
		return true
	} else if p.posX == selectedSquare.x && (p.posY+1 == selectedSquare.y || p.posY-1 == selectedSquare.y) { // Vertical move
		return true
	} else if math.Abs(float64(p.posX-selectedSquare.x)) == 1 && math.Abs(float64(p.posY-selectedSquare.y)) == 1 { // Diagonal move
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
