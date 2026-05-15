package handlers

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetJobs(c fiber.Ctx) error {
	return c.JSON(store.JobsData)
}
