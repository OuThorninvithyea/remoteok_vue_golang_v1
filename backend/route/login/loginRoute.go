package login

import (
	handlers "backend/handlers/login"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/data", handlers.GetLogin)
}
