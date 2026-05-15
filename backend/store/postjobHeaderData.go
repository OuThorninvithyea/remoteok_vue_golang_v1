package store

import "backend/model/postJob"

var PostjobHeaderData = postJob.PostJobHeader{
	Logo: postJob.PostJobLogoMark{
		Text:      "remote",
		Badge:     "OK",
		Trademark: "®",
		Link:      "/",
	},
	Title: "Hire remotely",
	Button: postJob.PostJobButton{
		Text: "Buy a bundle →",
		Link: "#",
	},
}
