package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetMainBar(c fiber.Ctx) error {
	return c.JSON(store.MainBarData)
}
