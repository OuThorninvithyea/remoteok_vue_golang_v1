package postJob

type PreviewCard struct {
	Logo                   string `json:"logo"`
	CompanyPlaceholder     string `json:"companyPlaceholder"`
	PositionPlaceholder    string `json:"positionPlaceholder"`
	LocationBadge          string `json:"locationBadge"`
	HiringText             string `json:"hiringText"`
	RoleText               string `json:"roleText"`
	DescriptionPlaceholder string `json:"descriptionPlaceholder"`
	ApplyHeading           string `json:"applyHeading"`
	ApplyPlaceholder       string `json:"applyPlaceholder"`
	ApplyButtonText        string `json:"applyButtonText"`
}

type PreviewSection struct {
	Legend     string      `json:"legend"`
	Heading    string      `json:"heading"`
	Subheading string      `json:"subheading"`
	Card       PreviewCard `json:"card"`
	Watermark  string      `json:"watermark"`
}
