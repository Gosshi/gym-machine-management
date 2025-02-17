package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gosshi/gym-machine-management/model"
	"github.com/matoous/go-nanoid/v2"
)

func CreateManufacturer(c *fiber.Ctx) error {
	var manufacturer model.Manufacturer

	// JSON をパース
	if err := c.BodyParser(&manufacturer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// ID を生成
	manufacturer.ID, _ = gonanoid.New()

	// DB に保存
	// result := config.DB.Create(&manufacturer)
	// if result.Error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Failed to create manufacturer",
	// 	})
	// }

	// 成功レスポンス
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": manufacturer.ID,
	})
}
