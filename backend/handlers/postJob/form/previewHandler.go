package postjobform

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetPreview(c fiber.Ctx) error {
	return c.JSON(store.PostjobPreviewData)
}
