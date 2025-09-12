package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Square struct {
	x int
	y int
}

const width = 800
const height = 800
const squareSize = 100

var board = [8][8]*Piece{
	{
		&Piece{0, 0, nil, 'R', true, false},
		&Piece{1, 0, nil, 'K', true, false},
		&Piece{2, 0, nil, 'B', true, false},
		&Piece{3, 0, nil, 'Q', true, false},
		&Piece{4, 0, nil, 'X', true, false},
		&Piece{5, 0, nil, 'B', true, false},
		&Piece{6, 0, nil, 'K', true, false},
		&Piece{7, 0, nil, 'R', true, false},
	},
	{
		&Piece{0, 1, nil, 'P', true, false},
		&Piece{1, 1, nil, 'P', true, false},
		&Piece{2, 1, nil, 'P', true, false},
		&Piece{3, 1, nil, 'P', true, false},
		&Piece{4, 1, nil, 'P', true, false},
		&Piece{5, 1, nil, 'P', true, false},
		&Piece{6, 1, nil, 'P', true, false},
		&Piece{7, 1, nil, 'P', true, false},
	},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{
		&Piece{0, 6, nil, 'P', false, false},
		&Piece{1, 6, nil, 'P', false, false},
		&Piece{2, 6, nil, 'P', false, false},
		&Piece{3, 6, nil, 'P', false, false},
		&Piece{4, 6, nil, 'P', false, false},
		&Piece{5, 6, nil, 'P', false, false},
		&Piece{6, 6, nil, 'P', false, false},
		&Piece{7, 6, nil, 'P', false, false},
	},
	{
		&Piece{0, 7, nil, 'R', false, false},
		&Piece{1, 7, nil, 'K', false, false},
		&Piece{2, 7, nil, 'B', false, false},
		&Piece{3, 7, nil, 'Q', false, false},
		&Piece{4, 7, nil, 'X', false, false},
		&Piece{5, 7, nil, 'B', false, false},
		&Piece{6, 7, nil, 'K', false, false},
		&Piece{7, 7, nil, 'R', false, false},
	},
}

func drawBoard(screen *ebiten.Image) {
	w := true
	for i, e := range board {
		for j, _ := range e {
			if w {
				vector.DrawFilledRect(screen, float32(j*squareSize), float32(i*squareSize), squareSize, squareSize,
					color.RGBA{234, 233, 210, 0}, false)
				w = false
			} else {
				vector.DrawFilledRect(screen, float32(j*squareSize), float32(i*squareSize), squareSize, squareSize,
					color.RGBA{75, 115, 153, 0}, false)
				w = true
			}
		}
		w = !w
	}
}
