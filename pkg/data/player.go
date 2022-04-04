package data

import "math/big"

func NewPlayer(name string, email string) *Player {
	return &Player{
		Name:  name,
		Email: email,
		Money: *big.NewFloat(0),
	}
}

// Player represents information about the human player
// and the character they play
type Player struct {
	Name  string
	Email string
	Money big.Float
}
