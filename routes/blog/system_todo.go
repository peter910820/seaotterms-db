package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	database "seaotterms-db/database/blog"
	blogdto "seaotterms-db/dto/blog"
	"seaotterms-db/middleware"
)

func systemTodoRouter(blogGroup fiber.Router, authDB *gorm.DB, db *gorm.DB) {
	systemTodoGroup := blogGroup.Group("/system-todos")

	systemTodoGroup.Post("/", middleware.GetTokenAccessLevel[blogdto.GetSystemTodo](authDB), func(c *fiber.Ctx) error {
		return database.QuerySystemTodo(c, db)
	})
}
