package blog

import (
	"time"

	"gorm.io/gorm"
)

// all user todos of the blog
type Todo struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Owner      string     `gorm:"NOT NULL" json:"owner"`
	Topic      string     `gorm:"NOT NULL" json:"topic"`
	Title      string     `gorm:"NOT NULL" json:"title"`
	Status     uint       `gorm:"NOT NULL" json:"status"`
	Deadline   *time.Time `json:"deadline"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	CreateName string     `gorm:"NOT NULL" json:"createName"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdateName string     `json:"updateName"`
}

// 查詢指定使用者的待辦事項
func FindTodosByOwner(db *gorm.DB, owner string) ([]Todo, error) {
	var data []Todo

	err := db.Where("owner = ?", owner).Order("created_at DESC").Find(&data).Error
	return data, err
}

// 新增待辦事項
func CreateTodo(db *gorm.DB, data *Todo) error {
	return db.Create(data).Error
}

// 更新待辦事項狀態
func UpdateTodoStatus(db *gorm.DB, id string, data *Todo) error {
	return db.Model(&Todo{}).Where("id = ?", id).
		Select("status", "updated_at", "update_name").
		Updates(data).Error
}

// 刪除待辦事項
func DeleteTodo(db *gorm.DB, id string) error {
	return db.Where("id = ?", id).Delete(&Todo{}).Error
}
