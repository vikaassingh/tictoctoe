package api

import (
	"fmt"
	"tictoctoe/model"
)

type StrategyService struct{}

type EasyBotPlayingStrategy struct {
	Type string
}

func (eazyBotPlayingStrategy EasyBotPlayingStrategy) GetType() string {
	return eazyBotPlayingStrategy.Type
}

func (eazyBotPlayingStrategy EasyBotPlayingStrategy) SuggestMove(board model.Board) *model.Cell {
	n := board.Size
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board.GetBoard()[i][j].GetCellState() == model.Empty {
				return &board.GetBoard()[i][j]
			}
		}
	}
	return nil
}

type RowWinningStrategy struct{}

func (rowWinningStrategy RowWinningStrategy) CheckWin(board model.Board, cell model.Cell) bool {
	currentPlayer := cell.GetPlayer()
	row := cell.GetRow()

	for i := 0; i < board.Size; i++ {
		currentCell := board.GetBoard()[row][i]
		if currentCell.GetCellState() == model.Empty || currentCell.GetPlayer() != currentPlayer {
			fmt.Println("row: ", row, "i: ", i)
			return false
		}
	}
	return true
}

type ColumnWinningStrategy struct{}

func (columnWinningStrategy ColumnWinningStrategy) CheckWin(board model.Board, cell model.Cell) bool {
	currentPlayer := cell.GetPlayer()
	column := cell.GetColumn()

	for i := 0; i < board.Size; i++ {
		currentCell := board.GetBoard()[i][column]
		if currentCell.GetCellState() == model.Empty || currentCell.GetPlayer() != currentPlayer {
			return false
		}
	}
	return true
}

type DiagonalWinningStrategy struct{}

func (diagonalWinningStrategy DiagonalWinningStrategy) CheckWin(board model.Board, cell model.Cell) bool {
	currentPlayer := cell.GetPlayer()
	n := board.Size
	// Check diagonal from top left to bottom right
	for i := 0; i < n; i++ {
		currentCell := board.GetBoard()[i][i]
		if currentCell.GetCellState() == model.Empty || currentCell.GetPlayer() != currentPlayer {
			return false
		}
	}
	return true
}

type AntiDiagonalWinningStrategy struct{}

func (diagonalWinningStrategy AntiDiagonalWinningStrategy) CheckWin(board model.Board, cell model.Cell) bool {
	currentPlayer := cell.GetPlayer()
	n := board.Size
	// Check diagonal from top right to bottom left
	for i := 0; i < n; i++ {
		currentCell := board.GetBoard()[i][n-i-1]
		if currentCell.GetCellState() == model.Empty || currentCell.GetPlayer() != currentPlayer {
			return false
		}
	}
	return true
}
