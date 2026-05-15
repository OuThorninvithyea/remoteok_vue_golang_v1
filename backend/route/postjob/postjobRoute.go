package postjob

import (
	postjobfooter "backend/handlers/postJob/footer"
	postjobform "backend/handlers/postJob/form"
	postjobheader "backend/handlers/postJob/header"
	postjobsidebar "backend/handlers/postJob/sidebar"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/header", postjobheader.GetHeader)
	r.Get("/promotion", postjobheader.GetPromotion)
	r.Get("/form", postjobform.GetForm)
	r.Get("/toolbar", postjobform.GetToolbar)
	r.Get("/let_start", postjobform.GetLetsStart)
	r.Get("/design_post", postjobform.GetDesignPost)
	r.Get("/job_details", postjobform.GetJobDetails)
	r.Get("/company", postjobform.GetCompany)
	r.Get("/feedback", postjobform.GetFeedback)
	r.Get("/preview", postjobform.GetPreview)
	r.Get("/partner", postjobform.GetPartner)
	r.Get("/sidebar", postjobsidebar.GetSidebar)
	r.Get("/footer", postjobfooter.GetFooter)
}
