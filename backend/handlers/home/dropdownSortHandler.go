package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetDropdownSort(c fiber.Ctx) error {
	return c.JSON(store.DropdownSortData)
}
