package store

import (
	"backend/model/home/dropdown"
	"backend/model/home/menu"
)

var DropdownLogoData = dropdown.DropdownLogo{
	Header: dropdown.DropdownLogoHeader{
		Join:  "🧑‍💻 Join Remote OK",
		Login: "👋 Log in",
	},
	Sections: []dropdown.DropdownLogoSection{
		{
			Title: "GENERAL",
			Items: []menu.MenuItem{
				{CategoryItem: "🏠 Frontpage", URL: "#"},
				{CategoryItem: "🏝️ Remote jobs", URL: "#"},
				{CategoryItem: "🌗 Dark mode", URL: "#"},
				{CategoryItem: "🧑‍💻 Hire remote workers", URL: "#"},
				{CategoryItem: "📮 Post a job", URL: "#"},
				{CategoryItem: "⭐️ Go premium", URL: "#"},
			},
		},
		{
			Title: "TOP JOBS",
			Items: []menu.MenuItem{
				{CategoryItem: "🤖 AI Jobs", URL: "#"},
				{CategoryItem: "⏰ Async jobs", URL: "#"},
				{CategoryItem: "🌎 Distributed team", URL: "#"},
				{CategoryItem: "🎧 Support jobs", URL: "#"},
				{CategoryItem: "🤓 Engineer jobs", URL: "#"},
				{CategoryItem: "🤓 Software jobs", URL: "#"},
				{CategoryItem: "👴 Senior jobs", URL: "#"},
				{CategoryItem: "🛠️ Technical jobs", URL: "#"},
				{CategoryItem: "💼 Management jobs", URL: "#"},
				{CategoryItem: "🚀 Growth jobs", URL: "#"},
				{CategoryItem: "🧑‍💼 Lead jobs", URL: "#"},
			},
		},
		{
			Title: "COMPANIES",
			Items: []menu.MenuItem{
				{CategoryItem: "📮 Post a remote job", URL: "#"},
				{CategoryItem: "📦 Buy a job bundle", URL: "#"},
				{CategoryItem: "🏷️ Ask for a discount", URL: "#"},
				{CategoryItem: "🛡 Health insurance for teams", URL: "#"},
				{CategoryItem: "🛡 Health insurance for noma...", URL: "#"},
			},
		},
		{
			Title: "FEEDS",
			Items: []menu.MenuItem{
				{CategoryItem: "🛠 Remote Jobs API", URL: "#"},
				{CategoryItem: "🪄 RSS feed", URL: "#"},
				{CategoryItem: "🪚 JSON feed", URL: "#"},
				{CategoryItem: "🟧 Hacker News mode", URL: "#"},
				{CategoryItem: "❌ Safe for work mode", URL: "#"},
			},
		},
		{
			Title: "HELP",
			Items: []menu.MenuItem{
				{CategoryItem: "💡 Ideas + bugs", URL: "#"},
				{CategoryItem: "🚀 Changelog", URL: "#"},
				{CategoryItem: "🛍️ Merch", URL: "#"},
				{CategoryItem: "🛟 FAQ & Help", URL: "#"},
			},
		},
	},
}
