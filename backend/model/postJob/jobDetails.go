package postJob

type LogoUpload struct {
	Label      string `json:"label"`
	ButtonText string `json:"buttonText"`
}

type BrandColorOption struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	Badge   string `json:"badge"`
	Checked bool   `json:"checked"`
	Price   int    `json:"price"`
}

type SalaryField struct {
	Label          string `json:"label"`
	Min            int    `json:"min"`
	Max            int    `json:"max"`
	Step           int    `json:"step"`
	MinPlaceholder string `json:"minPlaceholder"`
	MaxPlaceholder string `json:"maxPlaceholder"`
	HelpText       string `json:"helpText"`
}

type ApplyField struct {
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	Required    bool   `json:"required"`
	HelpText    string `json:"helpText"`
}

type JobDetails struct {
	Legend     string           `json:"legend"`
	LogoUpload LogoUpload       `json:"logoUpload"`
	BrandColor BrandColorOption `json:"brandColor"`
	Salary     SalaryField      `json:"salary"`
	Benefits   []string         `json:"benefits"`
	ApplyUrl   ApplyField       `json:"applyUrl"`
	ApplyEmail ApplyField       `json:"applyEmail"`
}
