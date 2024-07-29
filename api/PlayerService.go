package api

import (
	"tictoctoe/model"
)

type PlayerService struct {
}

func (playerService *PlayerService) CreateHumanPlayer(name string, symbol string) model.IPlayer {
	humanPlayer := &model.HumanPlayer{}
	humanPlayer = humanPlayer.SetName(name).SetSymbol(symbol).SetPlayerType(model.HUMAN)
	return humanPlayer
}

func (playerService *PlayerService) CreateBotPlayer(name string, symbol string, level int) model.IPlayer {
	var botPlayer *model.BotPlayer
	botPlayer = &model.BotPlayer{}
	botPlayingStrategy := EasyBotPlayingStrategy{Type: "Easy"}
	botPlayer = botPlayer.SetName(name).SetSymbol(symbol).SetPlayerType(model.BOT).SetBotDifficultyLevel(model.BotDifficultyLevel(level)).SetBotPlayingStrategy(botPlayingStrategy)
	return botPlayer

}
