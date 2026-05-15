package postJob

type CompanyField struct {
	Label         string `json:"label"`
	Type          string `json:"type"`
	Placeholder   string `json:"placeholder"`
	Required      bool   `json:"required"`
	HelpText      string `json:"helpText"`
	TextareaClass string `json:"textareaClass"`
}

type ShowLogoCheckbox struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
	Price   int    `json:"price"`
}

type Company struct {
	Legend           string           `json:"legend"`
	Fields           []CompanyField   `json:"fields"`
	PayLaterLabel    string           `json:"payLaterLabel"`
	ShowLogoCheckbox ShowLogoCheckbox `json:"showLogoCheckbox"`
}
