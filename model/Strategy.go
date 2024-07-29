package model

type Strategy struct{}

type IWinningStrategy interface {
	CheckWin(Board, Cell) bool
}

type IBotPlayingStrategy interface {
	SuggestMove(Board) *Cell
	GetType() string
}
