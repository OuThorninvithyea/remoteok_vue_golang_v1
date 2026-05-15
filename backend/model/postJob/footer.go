package postJob

type PostJobFooter struct {
	Logo                string `json:"logo"`
	CompanyPlaceholder  string `json:"companyPlaceholder"`
	PositionPlaceholder string `json:"positionPlaceholder"`
	LocationBadge       string `json:"locationBadge"`
	ApplyText           string `json:"applyText"`
	PostButtonText      string `json:"postButtonText"`
	BasePrice           int    `json:"basePrice"`
}
