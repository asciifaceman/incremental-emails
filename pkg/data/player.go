package data

import (
	"fmt"
	"math/big"
)

func NewPlayer(name string, email string) *Player {
	return &Player{
		Name:  name,
		Email: email,
		Money: *big.NewFloat(0),
		Inbox: &Inbox{
			Emails:  NewEmailContainer(),
			Folders: make(map[string]*EmailContainer),
		},
	}
}

// Player represents information about the human player
// and the character they play
type Player struct {
	Name  string
	Email string
	Money big.Float
	Inbox *Inbox
}

// NewEmailContainer returns a new abstract email storage
func NewEmailContainer() *EmailContainer {
	return &EmailContainer{
		Emails: *big.NewInt(0),
	}
}

// EmailContainer is an abstraction of a container holding emails
type EmailContainer struct {
	Emails big.Int
}

// Inbox represents the email data structure
type Inbox struct {
	Emails  *EmailContainer
	Folders map[string]*EmailContainer
}

// NewFolder creates a new named email container that is not the inbox
func (i *Inbox) NewFolder(name string) error {
	if i.Folders[name] != nil {
		return fmt.Errorf("the folder name [%s] already exists", name)
	}

	f := NewEmailContainer()
	i.Folders[name] = f
	return nil
}
