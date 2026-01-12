package discordbot

import "time"

type Member struct {
	UserID     string    `gorm:"primaryKey"`
	ServerID   string    `gorm:"primaryKey"`
	UserName   string    `gorm:"not null"`
	Level      uint      `gorm:"not null;default:1"`
	Exp        uint      `gorm:"not null;default:0"` // 該等級的經驗值，加上LevelUpExp才是該成員的所有經驗值
	LevelUpExp uint      `gorm:"not null;default:5"`
	JoinAt     time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}

type Log struct {
	LogID     string `gorm:"primaryKey"`
	ServerID  string `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	Message   string `gorm:"not null"`
	UpdatedAt time.Time
}

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
