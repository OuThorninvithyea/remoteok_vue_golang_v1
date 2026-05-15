package logo

import "backend/model/home/menu"

type LogoEntry struct {
	Join  string       `json:"join,omitempty"`
	Login string       `json:"login,omitempty"`
	URL   string       `json:"url,omitempty"`
	Type  string       `json:"type,omitempty"`
	Title string       `json:"title,omitempty"`
	Items []menu.MenuItem `json:"items,omitempty"`
}
