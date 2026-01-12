package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	blogmodel "seaotterms-db/blog"
	"seaotterms-db/dto"
	dtoblog "seaotterms-db/dto/blog"
	utils "seaotterms-db/utils"
)

func QuerySystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	requestData, ok := c.Locals("request_data").(dto.CommonRequest[dtoblog.GetSystemTodo])
	if !ok {
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, "request data not found", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	var err error
	var data []blogmodel.SystemTodo
	if requestData.Data.ID == nil && requestData.Data.SystemName == nil && requestData.Data.Status == nil {
		err = db.Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	} else {
		if requestData.Data.ID != nil {
			err = db.Where("id = ?", requestData.Data.ID).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
		} else {
			if requestData.Data.SystemName != nil {
				db = db.Where("system_name = ?", requestData.Data.SystemName)
			}
			if requestData.Data.Status != nil {
				db = db.Where("status = ?", requestData.Data.Status)
			}
			err = db.Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
		}
	}
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "找不到SystemTodo資料", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}
	logrus.Info("查詢SystemTodo資料成功")
	response := utils.ResponseFactory(c, fiber.StatusOK, "", &data)
	return c.Status(fiber.StatusOK).JSON(response)
}
