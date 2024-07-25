package model

type Cell struct {
	Row       int
	Column    int
	CellState CellState
	Player    Player
}

type CellState int

const (
	Empty CellState = iota
	OCCUPIED
)

func NewCell(row, column int) *Cell {
	return &Cell{Row: row, Column: column, CellState: Empty}
}

func (cell *Cell) GetRow() int {
	return cell.Row
}

func (cell *Cell) GetColumn() int {
	return cell.Column
}

func (cell *Cell) GetCellState() CellState {
	return cell.CellState
}

func (cell *Cell) GetPlayer() Player {
	return cell.Player
}

func (cell *Cell) SetRow(Row int) *Cell {
	cell.Row = Row
	return cell
}

func (cell *Cell) SetColumn(Column int) *Cell {
	cell.Column = Column
	return cell
}

func (cell *Cell) SetCellState(CellState CellState) *Cell {
	cell.CellState = CellState
	return cell
}

func (cell *Cell) SetPlayer(Player Player) *Cell {
	cell.Player = Player
	return cell
}
