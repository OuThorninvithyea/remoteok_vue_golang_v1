package postjobform

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetCompany(c fiber.Ctx) error {
	return c.JSON(store.PostjobCompanyData)
}
