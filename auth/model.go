package auth

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	AccessLevel string    `gorm:"NOT NULL" json:"accessLevel"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy   string    `gorm:"NOT NULL" json:"createBy"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

// seaotterms related service user
type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(50);NOT NULL;uniqueIndex" json:"username"`
	Password  string    `gorm:"type:varchar(64);NOT NULL" json:"-"`
	Email     string    `gorm:"type:varchar(100);NOT NULL;uniqueIndex" json:"email"`
	Avatar    string    `gorm:"type:varchar(255);NOT NULL;default:''" json:"avatar"`
	Exp       int       `gorm:"default:0" json:"exp"`
	IsAdmin   bool      `gorm:"default:false" json:"isAdmin"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	// soft delete
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
