package store

import "backend/model/home/category"

var Categories = []category.Category{
	{
		HoverText: "Suggested filter, tap to add",
		Filters: []category.Filter{
			{Emoji: "🎧", Label: "Support"},
			{Emoji: "🤓", Label: "Engineer"},
			{Emoji: "🤓", Label: "Software"},
			{Emoji: "👵", Label: "Senior"},
			{Emoji: "🛠", Label: "Technical"},
			{Emoji: "💼", Label: "Management"},
			{Emoji: "🚀", Label: "Growth"},
			{Emoji: "👩‍✈️", Label: "Lead"},
		},
	},
}
