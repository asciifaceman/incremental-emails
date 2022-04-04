package data

func NewPlayer(name string, email string) *Player {
	return &Player{
		Name:  name,
		Email: email,
	}
}

// Player represents information about the human player
// and the character they play
type Player struct {
	Name  string
	Email string
}
