package entities

import "time"

type UserActivityLog struct {
	UserID   uint
	Name     string
	LastSeen time.Time
}
