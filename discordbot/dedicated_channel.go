package discordbot

import (
	"time"
)

// Discord機器人通用專用頻道
//
// 存放特殊用途的頻道ID
type DedicatedChannel struct {
	ID        string `gorm:"primaryKey"`
	ServerID  string `gorm:"not null"`
	ChannelID string `gorm:"not null"`
	Type      string `gorm:"not null"`
	Priority  int    `gorm:"not null"`
	Program   int    `gorm:"not null;default:999"` // 專用的專案/程式，999是全部通用
	UpdatedAt time.Time
}
