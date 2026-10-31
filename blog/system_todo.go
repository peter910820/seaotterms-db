package blog

import "gorm.io/gorm"

// 查詢所有系統待辦事項
func FindAllSystemTodos(db *gorm.DB) ([]SystemTodo, error) {
	var data []SystemTodo

	err := db.Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 依 ID 查詢系統待辦事項
func FindSystemTodoByID(db *gorm.DB, id string) ([]SystemTodo, error) {
	var data []SystemTodo

	err := db.Where("id = ?", id).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 依系統名稱查詢系統待辦事項
func FindSystemTodosBySystemName(db *gorm.DB, systemName string) ([]SystemTodo, error) {
	var data []SystemTodo

	err := db.Where("system_name = ?", systemName).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 依狀態查詢系統待辦事項
func FindSystemTodosByStatus(db *gorm.DB, status int) ([]SystemTodo, error) {
	var data []SystemTodo

	err := db.Where("status = ?", status).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 依系統名稱及狀態查詢系統待辦事項
func FindSystemTodosBySystemNameAndStatus(db *gorm.DB, systemName string, status int) ([]SystemTodo, error) {
	var data []SystemTodo

	err := db.Where("system_name = ? AND status = ?", systemName, status).
		Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	return data, err
}

// 新增系統待辦事項
func CreateSystemTodo(db *gorm.DB, data *SystemTodo) error {
	return db.Create(data).Error
}

// 更新系統待辦事項
func UpdateSystemTodo(db *gorm.DB, id string, data *SystemTodo) error {
	return db.Model(&SystemTodo{}).Where("id = ?", id).
		Select("system_name", "title", "detail", "status", "deadline", "urgency", "updated_at", "updated_name").
		Updates(data).Error
}

// 更新系統待辦事項狀態
func QuickUpdateSystemTodo(db *gorm.DB, id string, data *SystemTodo) error {
	return db.Model(&SystemTodo{}).Where("id = ?", id).
		Select("status", "updated_at", "updated_name").
		Updates(data).Error
}

// 刪除系統待辦事項
func DeleteSystemTodo(db *gorm.DB, id string) error {
	return db.Where("id = ?", id).Delete(&SystemTodo{}).Error
}
