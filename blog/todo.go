package blog

import "gorm.io/gorm"

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
