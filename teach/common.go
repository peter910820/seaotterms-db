package teach

import "gorm.io/gorm"

// 定義一個私有的 scope，確保排序邏輯統一
func latestOrder(db *gorm.DB) *gorm.DB {
	return db.Order("COALESCE(updated_at, created_at) DESC")
}
