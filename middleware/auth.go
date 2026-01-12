package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"seaotterms-db/auth"
	"seaotterms-db/dto"
	"seaotterms-db/utils"
)

func GetTokenAccessLevel[T any](db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestData dto.CommonRequest[T]
		if err := c.BodyParser(&requestData); err != nil {
			logrus.Error(err)
			response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		var responseData auth.Token
		err := db.Where("id = ?", requestData.Token).First(&responseData).Error
		if err != nil {
			logrus.Error(err)
			response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		c.Locals("request_data", requestData)
		return c.Next()
	}
}
