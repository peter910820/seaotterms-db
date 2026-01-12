package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"seaotterms-db/auth"
	"seaotterms-db/blog"

	"gorm.io/gorm"
)

func Migration(dbName string, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file error: %v", err)
	}

	switch dbName {
	case os.Getenv("AUTH_DB_NAME"):
		db.AutoMigrate(&auth.Token{})
	case os.Getenv("DB_NAME"):
		db.AutoMigrate(&blog.User{})
		db.AutoMigrate(&blog.Tag{})
		db.AutoMigrate(&blog.Article{})
		db.AutoMigrate(&blog.Todo{})
		db.AutoMigrate(&blog.TodoTopic{})
		db.AutoMigrate(&blog.SystemTodo{})
	default:
		logrus.Fatal("error in migration function")
	}
}
