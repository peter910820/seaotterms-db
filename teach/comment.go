package teach

import (
	"time"
)

type Comment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Content     string    `gorm:"NOT NULL" json:"content"`
	ArticleID   uint      `gorm:"NOT NULL" json:"articleId"` // article table id
	CreatedAt   time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedName string    `json:"updatedName"`
}
