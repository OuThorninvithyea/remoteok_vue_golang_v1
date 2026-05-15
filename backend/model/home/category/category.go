package category

type Filter struct {
	Emoji string `json:"emoji"`
	Label string `json:"label"`
}

type Category struct {
	HoverText string   `json:"hoverText"`
	Filters   []Filter `json:"filters"`
}
