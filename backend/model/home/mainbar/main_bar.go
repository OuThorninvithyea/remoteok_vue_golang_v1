package mainbar

type MainBar struct {
	Emoji         string `json:"emoji"`
	Text          string `json:"text"`
	HighlightText string `json:"highlightText"`
	HighlightLink string `json:"highlightLink"`
	Suffix        string `json:"suffix"`
}
