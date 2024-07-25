package main

import (
	"fmt"
	"tictoctoe/model"
)

func main() {
	fmt.Println("Hello, playground")
	player := new(model.Player)
	player.SetPlayerType(model.HUMAN)

	game := model.NewGame(3, []model.Player{*player, *player}, nil)
	game.Board.PrintBoard()
}
