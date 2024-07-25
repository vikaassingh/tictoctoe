package model

type HumanPlayer struct {
	Player
	Age int
}

func (humanplayer *HumanPlayer) GetAge() int {
	return humanplayer.Age
}

func (humanplayer *HumanPlayer) SetAge(Age int) *HumanPlayer {
	humanplayer.Age = Age
	return humanplayer
}
