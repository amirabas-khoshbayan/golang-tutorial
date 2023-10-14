package main

import (
	gamepkg "golang-tutorial/game"
)

func main() {
	//Make players
	dog := gamepkg.NewDog("Axlie")
	cat := gamepkg.NewCat("Binky")
	owl := gamepkg.NewOwl("Hanna")
	//Make the game
	game := gamepkg.NewGame(50)

	//Join player to the game
	game.Join(dog).Join(cat).Join(owl)

	// Loop til there is a winner
	var winner gamepkg.Player
	for winner == nil {
		//Move players
		game.MovePlayers()
		//Check winner
		winner = game.CheckWinner()
		//Print the game
		game.Print()
	}
}
