package auth

import "time"

type Token struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	AccessLevel string    `gorm:"NOT NULL" json:"accessLevel"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy   string    `gorm:"NOT NULL" json:"createBy"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
