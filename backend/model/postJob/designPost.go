package postJob

type DesignCheckbox struct {
	Name          string `json:"name"`
	Text          string `json:"text"`
	Badge         string `json:"badge"`
	Checked       bool   `json:"checked"`
	Price         int    `json:"price"`
	LinkText      string `json:"linkText"`
	LinkUrl       string `json:"linkUrl"`
	AfterLinkText string `json:"afterLinkText"`
}

type DesignPost struct {
	Legend                   string           `json:"legend"`
	Checkboxes               []DesignCheckbox `json:"checkboxes"`
	ViewCounterPlaceholder   string           `json:"viewCounterPlaceholder"`
	ClicksCounterPlaceholder string           `json:"clicksCounterPlaceholder"`
	Disclaimer               string           `json:"disclaimer"`
}
