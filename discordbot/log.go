package discordbot

import (
	"time"
)

type Log struct {
	LogID     string `gorm:"primaryKey"`
	ServerID  string `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	Message   string `gorm:"not null"`
	UpdatedAt time.Time
}
