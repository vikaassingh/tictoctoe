package model

type Game struct {
	Board              *Board
	Players            []IPlayer
	Moves              []Move
	WinningStrategies  []IWinningStrategy
	CurrentPlayerIndex int
	GameStatus         GameStatus
	Winner             IPlayer
}

func NewGame(dimension int, players []IPlayer, winningStrategies []IWinningStrategy) *Game {
	brd := Board{}
	return &Game{
		Board:              brd.SetSize(dimension).CreateBoard(),
		Players:            players,
		Moves:              []Move{},
		WinningStrategies:  winningStrategies,
		CurrentPlayerIndex: 0,
		GameStatus:         INITIAL,
		Winner:             nil,
	}
}
