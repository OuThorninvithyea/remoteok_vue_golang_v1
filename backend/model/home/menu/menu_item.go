package menu

type MenuItem struct {
	CategoryItem string `json:"categoryItem"`
	URL          string `json:"url"`
	Default      bool   `json:"default,omitempty"`
}
