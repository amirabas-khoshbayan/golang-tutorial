package game

import (
	"fmt"
	"golang-tutorial/tools"
	"time"
)

func main() {
	//Make players
	dog := NewDog("Axlie")
	cat := NewCat("Binky")
	owl := NewOwl("Hanna")

	//Make the game
	game := NewGame(50)

	//Join player to the game
	game.Join(dog).Join(cat).Join(owl)

	// Loop til there is a winner
	var winner Player
	for winner == nil {
		//Move players
		game.MovePlayers()
		//Check winner
		winner = game.CheckWinner()
		//Print the game
		tools.ClearTerminal()
		game.Print()
		time.Sleep(time.Millisecond * 250)
	}

	//Print the winner
	fmt.Printf("%s won the game!", winner)
}
