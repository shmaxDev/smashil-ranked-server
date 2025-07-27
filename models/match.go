package models

import "time"

type Match struct {
	ID           int
	Player1ID    string
	Player2ID    string
	WinnerID     string
	Player1Score int
	Player2Score int
	CreatedAt    time.Time
}
