package model

type BotPlayer struct {
	Player
	BotDifficultyLevel BotDifficultyLevel
}

func (botPlayer *BotPlayer) GetBotDifficultyLevel() BotDifficultyLevel {
	return botPlayer.BotDifficultyLevel
}

func (botPlayer *BotPlayer) SetBotDifficultyLevel(BotDifficultyLevel BotDifficultyLevel) *BotPlayer {
	botPlayer.BotDifficultyLevel = BotDifficultyLevel
	return botPlayer
}
