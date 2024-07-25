package model

type Game struct {
	Board              Board
	Players            []Player
	Moves              []Move
	WinningStrategies  []WinningStrategy
	CurrentPlayerIndex int
	GameStatus         GameStatus
	Winner             Player
}

func NewGame(dimension int, players []Player, winningStrategies []WinningStrategy) *Game {
	brd := Board{}
	return &Game{
		Board:              *brd.SetSize(dimension).CreateBoard(),
		Players:            players,
		Moves:              []Move{},
		WinningStrategies:  winningStrategies,
		CurrentPlayerIndex: 0,
		GameStatus:         INITIAL,
		// Winner: nil,

	}
}
