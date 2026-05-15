package postjobfooter

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetFooter(c fiber.Ctx) error {
	return c.JSON(store.PostjobFooterData)
}
