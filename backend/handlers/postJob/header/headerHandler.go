package postjobheader

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetHeader(c fiber.Ctx) error {
	return c.JSON(store.PostjobHeaderData)
}
