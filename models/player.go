package models

import "time"

type Player struct {
	DiscordID  int64
	DiscordTag string
	Elo        int
	CreatedAt  time.Time
}