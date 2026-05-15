package route

import (
	"backend/route/home"
	"backend/route/login"
	"backend/route/postjob"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	home.RegisterRoutes(api.Group("/home"))
	login.RegisterRoutes(api.Group("/login"))
	postjob.RegisterRoutes(api.Group("/postjob"))
}
