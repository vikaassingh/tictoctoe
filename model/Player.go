package model

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

func (player *Player) GetID() int {
	return player.ID
}

func (player *Player) GetRank() int {
	return player.Rank
}

func (player *Player) GetName() string {
	return player.Name
}

func (player *Player) GetSymbol() string {
	return player.Symbol
}

func (player *Player) GetPlayerType() PlayerType {
	return player.PlayerType
}

func (player *Player) SetID(ID int) *Player {
	player.ID = ID
	return player
}

func (player *Player) SetRank(Rank int) *Player {
	player.Rank = Rank
	return player
}

func (player *Player) SetName(Name string) *Player {
	player.Name = Name
	return player
}

func (player *Player) SetSymbol(Symbol string) *Player {
	player.Symbol = Symbol
	return player
}

func (player *Player) SetPlayerType(PlayerType PlayerType) *Player {
	player.PlayerType = PlayerType
	return player
}
