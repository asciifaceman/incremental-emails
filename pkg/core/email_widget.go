package core

import "time"

/*
	An email looks like this

	Subject content     {optional attachment icon}     Date Received

	3 labels spaced out, all tappable as one
*/

type EmailWidget struct {
	Subject       string
	HasAttachment bool
	Received      time.Time
}
