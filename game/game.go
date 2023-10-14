package game

type Player interface {
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
