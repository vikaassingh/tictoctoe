package model

type Move struct {
	Cell   Cell
	Player Player
}

func (move *Move) GetCell() Cell {
	return move.Cell
}

func (move *Move) GetPlayer() Player {
	return move.Player
}

func (move *Move) SetCell(Cell Cell) *Move {
	move.Cell = Cell
	return move
}

func (move *Move) SetPlayer(Player Player) *Move {
	move.Player = Player
	return move
}
