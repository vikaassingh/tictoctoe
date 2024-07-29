package model

import "fmt"

type HumanPlayer struct {
	Player
	Age int
}

func (player HumanPlayer) GetID() int {
	return player.ID
}

func (player HumanPlayer) GetRank() int {
	return player.Rank
}

func (player HumanPlayer) GetName() string {
	return player.Name
}

func (player HumanPlayer) GetSymbol() string {
	return player.Symbol
}

func (player HumanPlayer) GetPlayerType() PlayerType {
	return player.PlayerType
}

func (player HumanPlayer) SetID(ID int) *HumanPlayer {
	player.ID = ID
	return &player
}

func (player HumanPlayer) SetRank(Rank int) *HumanPlayer {
	player.Rank = Rank
	return &player
}

func (player HumanPlayer) SetName(name string) *HumanPlayer {
	player.Name = name
	return &player
}

func (player HumanPlayer) SetSymbol(Symbol string) *HumanPlayer {
	player.Symbol = Symbol
	return &player
}

func (player HumanPlayer) SetPlayerType(PlayerType PlayerType) *HumanPlayer {
	player.PlayerType = PlayerType
	return &player
}

func (humanPlayer HumanPlayer) GetAge() int {
	return humanPlayer.Age
}

func (humanPlayer HumanPlayer) SetAge(Age int) *HumanPlayer {
	humanPlayer.Age = Age
	return &humanPlayer
}

func (humanPlayer HumanPlayer) NextMove(board Board) *Cell {
	var row, column int
	fmt.Println("Please enter the row")
	fmt.Scan(&row)
	fmt.Println("Please enter the column")
	fmt.Scan(&column)
	if board.GetBoard()[row][column].GetCellState() != Empty {
		fmt.Println("Cell is not empty")
	}
	cell := &board.GetBoard()[row][column]
	cell.SetPlayer(humanPlayer)
	cell.SetCellState(OCCUPIED)
	return cell
}
