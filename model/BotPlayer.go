package model

import "fmt"

type BotPlayer struct {
	Player
	BotDifficultyLevel BotDifficultyLevel
	BotPlayingStrategy IBotPlayingStrategy
}

func (botPlayer BotPlayer) GetID() int {
	return botPlayer.ID
}

func (botPlayer BotPlayer) GetRank() int {
	return botPlayer.Rank
}

func (botPlayer BotPlayer) GetName() string {
	return botPlayer.Name
}

func (botPlayer BotPlayer) GetSymbol() string {
	return botPlayer.Symbol
}

func (botPlayer BotPlayer) GetPlayerType() PlayerType {
	return botPlayer.PlayerType
}

func (botPlayer BotPlayer) SetID(ID int) *BotPlayer {
	botPlayer.ID = ID
	return &botPlayer
}

func (botPlayer BotPlayer) SetRank(Rank int) *BotPlayer {
	botPlayer.Rank = Rank
	return &botPlayer
}

func (botPlayer BotPlayer) SetName(Name string) *BotPlayer {
	botPlayer.Name = Name
	return &botPlayer
}

func (botPlayer BotPlayer) SetSymbol(Symbol string) *BotPlayer {
	botPlayer.Symbol = Symbol
	return &botPlayer
}

func (botPlayer BotPlayer) SetPlayerType(PlayerType PlayerType) *BotPlayer {
	botPlayer.PlayerType = PlayerType
	return &botPlayer
}

func (botPlayer BotPlayer) GetBotDifficultyLevel() BotDifficultyLevel {
	return botPlayer.BotDifficultyLevel
}

func (botPlayer BotPlayer) SetBotDifficultyLevel(BotDifficultyLevel BotDifficultyLevel) *BotPlayer {
	botPlayer.BotDifficultyLevel = BotDifficultyLevel
	return &botPlayer
}

func (botPlayer BotPlayer) GetBotPlayingStrategy() IBotPlayingStrategy {
	return botPlayer.BotPlayingStrategy
}

func (botPlayer BotPlayer) SetBotPlayingStrategy(botPlayingStrategy IBotPlayingStrategy) *BotPlayer {
	botPlayer.BotPlayingStrategy = botPlayingStrategy
	fmt.Println("Bot Playing strategy is set:", botPlayer.BotPlayingStrategy.GetType())
	return &botPlayer
}

func (botPlayer BotPlayer) NextMove(board Board) *Cell {
	botPlayerStrategy := botPlayer.GetBotPlayingStrategy()
	cell := botPlayerStrategy.SuggestMove(board)
	cell.SetPlayer(botPlayer)
	cell.SetCellState(OCCUPIED)
	// boardCell := board.GetBoard()[cell.GetRow()][cell.GetColumn()]
	// boardCell.SetPlayer(botPlayer)
	// boardCell.SetCellState(OCCUPIED)
	return cell
}
