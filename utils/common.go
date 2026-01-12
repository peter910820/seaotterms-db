package utils

import (
	"github.com/gofiber/fiber/v2"

	"seaotterms-db/dto"
)

func ResponseFactory[T any](c *fiber.Ctx, httpStatus int, msg string, data T) dto.CommonResponse[T] {
	return dto.CommonResponse[T]{
		StatusCode: httpStatus,
		Data:       data,
		ErrMsg:     msg,
	}
}
