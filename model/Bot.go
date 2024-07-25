package model

type Bot struct {
	BotDifficultyLevel
}

type BotDifficultyLevel int

const (
	Easy BotDifficultyLevel = iota
	Medium
	Hard
)
