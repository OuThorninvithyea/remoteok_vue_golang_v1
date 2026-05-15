package postJob

type FormField struct {
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	Required    bool   `json:"required"`
	HelpText    string `json:"helpText"`
}

type FormFieldWithNew struct {
	Label         string `json:"label"`
	IsNew         bool   `json:"isNew"`
	Placeholder   string `json:"placeholder"`
	Required      bool   `json:"required"`
	HelpText      string `json:"helpText"`
	HighlightText string `json:"highlightText"`
	UnderlineText string `json:"underlineText"`
}

type SelectOption struct {
	Value    string `json:"value"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled,omitempty"`
}

type JobDescriptionField struct {
	Label        string `json:"label"`
	AIButtonText string `json:"aiButtonText"`
}

type TagsField struct {
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	HelpText    string `json:"helpText"`
}

type LetsStart struct {
	Legend             string              `json:"legend"`
	Position           FormField           `json:"position"`
	CompanyName        FormField           `json:"companyName"`
	JobDescription     JobDescriptionField `json:"jobDescription"`
	EmploymentTypes    []SelectOption      `json:"employmentTypes"`
	PrimaryTags        []SelectOption      `json:"primaryTags"`
	PrimaryTagHelpText string              `json:"primaryTagHelpText"`
	Tags               TagsField           `json:"tags"`
	Location           FormFieldWithNew    `json:"location"`
}
