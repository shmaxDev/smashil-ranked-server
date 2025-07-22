package models

import "time"

type Player struct {
	DiscordID  string
	DiscordUsername string
	Elo        int
	CreatedAt  time.Time
}