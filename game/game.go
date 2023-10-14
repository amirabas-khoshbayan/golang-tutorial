package game

import (
	"fmt"
	"strings"
)

type Player interface {
	Move()
	Position() int
	Name() string
}

type Game struct {
	fieldLength int
	players     []Player
}

func NewGame(fieldsLenghth int) *Game {
	return &Game{
		fieldLength: fieldsLenghth,
		players:     make([]Player, 0),
	}
}
func (g *Game) Join(player Player) *Game {
	g.players = append(g.players, player)
	return g
}

func (g *Game) MovePlayers() {
	for _, player := range g.players {
		player.Move()
	}
}
func (g *Game) CheckWinner() Player {

	for _, player := range g.players {
		if player.Position() > g.fieldLength {
			return player
		}
	}
	return nil
}

func (g *Game) Print() {
	fmt.Println("|" + strings.Repeat("-", g.fieldLength+6) + "|")
	for _, player := range g.players {
		pos := player.Position()
		name := player.Name()
		repeat_count := g.fieldLength - pos + 1
		if repeat_count < 0 {
			repeat_count = 0
		}
		row := "|" + strings.Repeat(" ", pos) + name + strings.Repeat(" ", repeat_count) + "|"
		fmt.Println(row)
	}
	fmt.Println("|" + strings.Repeat("-", g.fieldLength+6) + "|")
}
