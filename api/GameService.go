package api

import (
	"fmt"
	"tictoctoe/model"
)

type GameService struct {
	Game *model.Game
}

func (gameService *GameService) InitiateGame(dimension int, players []model.IPlayer, winningStrategies []model.IWinningStrategy) *model.Game {
	return model.NewGame(dimension, players, winningStrategies)
}

func (gameService *GameService) StartGame() {
	gameService.Game.GameStatus = model.IN_PROGRESS
	gameService.ExecuteNextMove()
}

func (gameService *GameService) CheckEmptyCell() bool {
	n := gameService.Game.Board.GetSize()
	board := gameService.Game.Board.GetBoard()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j].GetCellState() == model.Empty {
				return true
			}
		}
	}
	return false
}

func (gameService *GameService) GetGame() *model.Game {
	return gameService.Game
}

func (gameService *GameService) SetGame(game *model.Game) *GameService {
	gameService.Game = game
	return gameService
}

func (gameService *GameService) ExecuteNextMove() {
	game := gameService.Game
	game.Board.PrintBoard()
	for gameService.CheckEmptyCell() {
		currentPlayer := game.Players[game.CurrentPlayerIndex]
		fmt.Println("It's " + currentPlayer.GetName() + " move")
		cell := currentPlayer.NextMove(*game.Board)
		game.Moves = append(game.Moves, model.Move{Player: currentPlayer, Cell: *cell})
		game.Board.PrintBoard()

		for _, strategy := range game.WinningStrategies {
			if strategy.CheckWin(*game.Board, *cell) {
				game.GameStatus = model.SUCCESS
				game.Winner = currentPlayer
				break
			}
		}

		if game.GameStatus == model.SUCCESS {
			if game.GameStatus == model.DRAW {
				fmt.Println("Game is draw")
			} else {
				fmt.Println("Winner is " + game.Winner.GetName())
			}
			break
		}
		game.CurrentPlayerIndex = (game.CurrentPlayerIndex + 1) % len(game.Players)
	}
}
