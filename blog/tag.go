package blog

import "gorm.io/gorm"

// 計算指定名稱的標籤數量
func CountTags(db *gorm.DB, names []string) (int64, error) {
	var count int64

	err := db.Model(&Tag{}).Where("name IN ?", names).Count(&count).Error
	return count, err
}

// 查詢所有標籤
func FindAllTags(db *gorm.DB) ([]Tag, error) {
	var data []Tag

	err := db.Order("created_at desc").Find(&data).Error
	return data, err
}

// 新增標籤
func CreateTag(db *gorm.DB, data *Tag) error {
	return db.Create(data).Error
}
