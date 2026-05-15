package postjobform

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetJobDetails(c fiber.Ctx) error {
	return c.JSON(store.PostjobJobDetailsData)
}
