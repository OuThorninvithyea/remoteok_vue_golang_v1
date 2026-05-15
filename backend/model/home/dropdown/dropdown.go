package dropdown

import "backend/model/home/menu"

type DropdownCategory struct {
	Data string `json:"data"`
}

type DropdownLogoHeader struct {
	Join  string `json:"join"`
	Login string `json:"login"`
}

type DropdownLogoSection struct {
	Title string          `json:"title"`
	Items []menu.MenuItem `json:"items"`
}

type DropdownLogo struct {
	Header   DropdownLogoHeader    `json:"header"`
	Sections []DropdownLogoSection `json:"sections"`
}

type LocationData struct {
	Regions   []menu.MenuItem `json:"regions"`
	Countries []menu.MenuItem `json:"countries"`
}
