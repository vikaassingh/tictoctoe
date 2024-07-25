package model

type WinningStrategy interface {
	CheckWin(Board, Move) bool
}

type BotPlayingStrategy interface {
	MakeMove(Board) Move
}
