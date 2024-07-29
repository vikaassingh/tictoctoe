package model

import "fmt"

type Board struct {
	Size  int
	Board [][]Cell
}

func (board *Board) CreateBoard() *Board {
	var brd [][]Cell
	for i := 0; i < board.Size; i++ {
		brd = append(brd, make([]Cell, board.Size))
		for j := 0; j < board.Size; j++ {
			brd[i][j] = Cell{Row: i, Column: j}
		}
	}
	board.Board = brd
	return board
}

func (board *Board) PrintBoard() {
	for i := 0; i < board.Size; i++ {
		fmt.Printf("|\t")
		for j := 0; j < board.Size; j++ {
			if board.Board[i][j].CellState == Empty {
				fmt.Printf(" \t|\t")
			} else {
				fmt.Printf(board.Board[i][j].GetPlayer().GetSymbol() + "\t|\t")
			}
		}
		fmt.Println()
	}
}

func (board *Board) GetSize() int {
	return board.Size
}

func (board *Board) GetBoard() [][]Cell {
	return board.Board
}

func (board *Board) SetSize(Size int) *Board {
	board.Size = Size
	return board
}

func (board *Board) SetBoard(Board [][]Cell) *Board {
	board.Board = Board
	return board
}
