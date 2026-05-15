package postjobheader

import (
	"backend/store"
	"github.com/gofiber/fiber/v3"
)

func GetPromotion(c fiber.Ctx) error {
	return c.JSON(store.PostjobPromotionData)
}
