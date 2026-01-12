package blog

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BlogRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	blogGroup := apiGroup.Group("/blog")
	dbName := os.Getenv("DB_NAME")

	systemTodoRouter(blogGroup, dbs[os.Getenv("AUTH_DB_NAME")], dbs[dbName])
}
