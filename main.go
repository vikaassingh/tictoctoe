package main

import (
	"fmt"
	api "tictoctoe/api"
	"tictoctoe/model"
)

var (
	playerService = api.PlayerService{}
	gameService   = api.GameService{}
)

func main() {
	fmt.Println("Welcome to the Tic Toc Toe game")
	fmt.Println("Let's start the game")
	var players []model.IPlayer
	fmt.Println("How many players are playing?")
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter the player name")
		var name string
		fmt.Scan(&name)
		fmt.Println("Enter the player symbol")
		var symbol string
		fmt.Scan(&symbol)
		fmt.Println("Enter the player type Bot or Human?(B/H)")
		var playerType string
		fmt.Scan(&playerType)
		var player model.IPlayer
		if playerType == "B" || playerType == "b" {
			var level int
			fmt.Println("What is the level of Bot? (0/1/2)")
			fmt.Scan(&level)
			player = playerService.CreateBotPlayer(name, symbol, level)
			players = append(players, player)
		} else {
			player = playerService.CreateHumanPlayer(name, symbol)
			players = append(players, player)
		}
		fmt.Println("Player name: ", player.GetName(), "Player symbol: ", player.GetSymbol(), "Player type: ", player.GetPlayerType())
	}

	game := gameService.InitiateGame(3, players, []model.IWinningStrategy{
		new(api.RowWinningStrategy),
		new(api.ColumnWinningStrategy),
		&api.DiagonalWinningStrategy{},
		&api.AntiDiagonalWinningStrategy{},
	})
	// game.Board.PrintBoard()
	fmt.Println("Are you ready to start the game? (Y/N)")
	var start string
	fmt.Scan(&start)
	fmt.Println("Game started")
	if start == "Y" || start == "y" {
		gameService.SetGame(game).StartGame()
	}
}
