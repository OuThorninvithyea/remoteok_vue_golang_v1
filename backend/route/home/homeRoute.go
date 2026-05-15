package home

import (
	"backend/handlers/home"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/categories", handlers.GetCategories)
	r.Get("/jobs", handlers.GetJobs)
	r.Get("/main-bar", handlers.GetMainBar)
	r.Get("/footer", handlers.GetFooter)
	r.Get("/logo", handlers.GetLogo)
	r.Get("/dropdown/search", handlers.GetDropdownSearch)
	r.Get("/dropdown/sort", handlers.GetDropdownSort)
	r.Get("/dropdown/location", handlers.GetDropdownLocation)
	r.Get("/dropdown/salary", handlers.GetDropdownSalary)
	r.Get("/dropdown/benefits", handlers.GetDropdownBenefits)
}
