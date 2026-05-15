package postjobsidebar

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetSidebar(c fiber.Ctx) error {
	return c.JSON(store.PostjobSidebarData)
}
