package models

type Game struct {
	Score int
}

func NewGame() *Game {
	return &Game{Score: 0}
}
