package blog

import (
	"time"

	"gorm.io/gorm"
)

// all user info of the blog
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"NOT NULL unique" json:"username"`
	Password     string    `gorm:"NOT NULL" json:"-"`
	Email        string    `gorm:"NOT NULL unique" json:"email"`
	Avatar       string    `gorm:"NOT NULL; default:''" json:"avatar"`
	Introduction string    `gorm:"NOT NULL; default:''" json:"introduction"`
	Exp          int       `gorm:"default:0" json:"exp"`
	Management   bool      `gorm:"default:false" json:"management"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreateName   string    `gorm:"NOT NULL" json:"createName"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdateName   string    `json:"updateName"`
}

// 查詢所有使用者(依更新時間排序)
func FindAllUsers(db *gorm.DB) ([]User, error) {
	var data []User

	err := db.Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 依使用者名稱查詢使用者
func FindUserByUsername(db *gorm.DB, username string) (*User, error) {
	var data User

	err := db.Where("username = ?", username).First(&data).Error
	return &data, err
}

// 新增使用者
func CreateUser(db *gorm.DB, data *User) error {
	return db.Create(data).Error
}

// 更新使用者個人資料
func UpdateUserProfile(db *gorm.DB, id uint64, data *User) error {
	return db.Model(&User{}).Where("id = ?", id).
		Select("updated_at", "update_name", "avatar", "introduction").
		Updates(data).Error
}

// 更新使用者管理權限
func UpdateUserManagement(db *gorm.DB, id uint64, data *User) error {
	return db.Model(&User{}).Where("id = ?", id).
		Select("updated_at", "update_name", "management").
		Updates(data).Error
}
