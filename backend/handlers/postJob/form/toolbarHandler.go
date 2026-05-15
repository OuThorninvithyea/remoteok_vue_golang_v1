package postjobform

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetToolbar(c fiber.Ctx) error {
	return c.JSON(store.PostjobToolbarData)
}
