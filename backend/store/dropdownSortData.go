package store

import "backend/model/home/menu"

var DropdownSortData = []menu.MenuItem{
	{CategoryItem: "🦴 Sort by", URL: "#", Default: true},
	{CategoryItem: "🆕 Latest jobs", URL: "#"},
	{CategoryItem: "💵 Highest paid", URL: "#"},
	{CategoryItem: "👀 Most viewed", URL: "#"},
	{CategoryItem: "✅ Most applied", URL: "#"},
	{CategoryItem: "🔥 Hottest", URL: "#"},
	{CategoryItem: "🎪 Most benefits", URL: "#"},
}
