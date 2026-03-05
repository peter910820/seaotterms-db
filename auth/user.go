package auth

import (
	"time"

	"gorm.io/gorm"
)

// retrieves all non-deleted users
func FindAllUsers(db *gorm.DB) ([]User, error) {
	var data []User
	err := db.Find(&data).Error
	return data, err
}

// retrieves a single user by primary key
func FindUserByID(db *gorm.DB, id int) (*User, error) {
	var data User
	err := db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// retrieves a single user by username
func FindUserByUsername(db *gorm.DB, username string) (*User, error) {
	var data User
	err := db.Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// retrieves a single user by email
func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	var data User
	err := db.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// inserts a new user record into the database
func CreateUser(db *gorm.DB, data *User) error {
	return db.Create(data).Error
}

// updates only the password field of the user specified by id.
func UpdateUserPassword(db *gorm.DB, id int, password string) error {
	return db.Model(&User{}).Where("id = ?", id).
		Select("password", "updated_at").
		Updates(map[string]any{
			"password":   password,
			"updated_at": time.Now(),
		}).Error
}

// soft-deletes the user specified by id
func DeleteUser(db *gorm.DB, id int) error {
	return db.Delete(&User{}, id).Error
}
