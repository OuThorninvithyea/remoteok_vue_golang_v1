package postjobform

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetFeedback(c fiber.Ctx) error {
	return c.JSON(store.PostjobFeedbackData)
}
