package store

import (
	"backend/model/home/dropdown"
	"backend/model/home/menu"
)

var DropdownLocationData = dropdown.LocationData{
	Regions: []menu.MenuItem{
		{CategoryItem: "🌍 Worldwide", URL: "#"},
		{CategoryItem: "⛰️ North America", URL: "#"},
		{CategoryItem: "💃 Latin America", URL: "#"},
		{CategoryItem: "🇪🇺 Europe", URL: "#"},
		{CategoryItem: "🦁 Africa", URL: "#"},
		{CategoryItem: "🕌 Middle East", URL: "#"},
		{CategoryItem: "⛩️ Asia", URL: "#"},
		{CategoryItem: "🌊 Oceania", URL: "#"},
	},
	Countries: []menu.MenuItem{
		{CategoryItem: "🇺🇸 United States", URL: "#"},
		{CategoryItem: "🇨🇦 Canada", URL: "#"},
		{CategoryItem: "🇬🇧 United Kingdom", URL: "#"},
		{CategoryItem: "🇦🇺 Australia", URL: "#"},
		{CategoryItem: "🇳🇿 New Zealand", URL: "#"},
	},
}
