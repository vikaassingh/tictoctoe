package model

type Move struct {
	Cell   Cell
	Player IPlayer
}

func (move *Move) GetCell() Cell {
	return move.Cell
}

func (move *Move) GetPlayer() IPlayer {
	return move.Player
}

func (move *Move) SetCell(Cell Cell) *Move {
	move.Cell = Cell
	return move
}

func (move Move) SetPlayer(Player IPlayer) *Move {
	move.Player = Player
	return &move
}
