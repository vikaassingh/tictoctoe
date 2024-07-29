package model

type IPlayer interface {
	NextMove(Board) *Cell
	GetID() int
	GetRank() int
	GetName() string
	GetSymbol() string
	GetPlayerType() PlayerType
}

type Player struct {
	ID         int
	Rank       int
	Name       string
	Symbol     string
	PlayerType PlayerType
}

type PlayerType int

const (
	HUMAN PlayerType = iota
	BOT
)

// func (p *Player) SetPlayerType(playerType PlayerType) {
// 	p.PlayerType = playerType
// }
