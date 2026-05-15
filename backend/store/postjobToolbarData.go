package store

import "backend/model/postJob"

var PostjobToolbarData = []postJob.ToolbarIconGroup{
	{Group: "undo-do", Icons: []string{"/images/icon1.svg", "/images/icon2.svg"}},
	{Group: "font-style", Icons: []string{"/images/icon3.svg", "/images/icon4.svg"}},
	{Group: "font-size", Icons: []string{"/images/icon5.svg", "/images/icon6.svg"}},
	{Group: "font-list", Icons: []string{"/images/icon7.svg", "/images/icon8.svg"}},
	{Group: "font-align", Icons: []string{"/images/icon9.svg", "/images/icon10.svg"}},
	{Group: "attach-folders", Icons: []string{"/images/icon11.svg"}},
}
