package model

type GameStatus int

const (
	IN_PROGRESS GameStatus = iota
	PAUSED
	DRAW
	SUCCESS
	INITIAL
)
