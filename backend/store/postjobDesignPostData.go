package store

import "backend/model/postJob"

var PostjobDesignPostData = postJob.DesignPost{
	Legend: "DESIGN YOUR JOB POST",
	Checkboxes: []postJob.DesignCheckbox{
		{Name: "show_logo", Text: "Show my ⭐️ company logo besides my post (+$44)", Badge: "2X MORE VIEWS", Checked: true, Price: 44},
		{Name: "email_blast", Text: "Email blast my job post to 📮 2,309,440 remote candidates (+$89)", Badge: "3X MORE VIEWS", Checked: true, Price: 89},
		{Name: "auto_match", Text: "Get auto matched with suitable applicants from our", LinkText: "directory of remote workers", LinkUrl: "#", Checked: true, Price: 0},
		{Name: "qr_code", Text: "Create a 🔲 QR-code short link for your post like", LinkText: "rok.co/qi", LinkUrl: "#", AfterLinkText: "to share easily (+$44)", Checked: true, Price: 44},
		{Name: "highlight_yellow", Text: "Highlight your post in ⚠️ yellow (+$49)", Badge: "2X MORE VIEWS", Checked: true, Price: 49},
		{Name: "sticky_24h", Text: "Sticky your post so it stays on 📌 top of the FrontPage for ⏰ 24 hours (+$89)", Badge: "2X MORE VIEWS", Checked: false, Price: 89},
		{Name: "sticky_week", Text: "Sticky your post so it stays on 📌 top of the FrontPage for 🗓 1 entire week (+$179)", Badge: "6X MORE VIEWS", Checked: false, Price: 179},
		{Name: "sticky_month", Text: "Sticky your post so it stays on 📌 top of the FrontPage for 🗓 1 entire month (+$537)", Badge: "9X MORE VIEWS", Checked: false, Price: 537},
		{Name: "geologic", Text: "Geologic your post for people only in the location it's restricted to above and block applicants from elsewhere (+$89)", Checked: false, Price: 89},
	},
	ViewCounterPlaceholder:   "2,272",
	ClicksCounterPlaceholder: "266",
	Disclaimer:               "Pricing of job posts and extra features is dynamic and may change based on how many jobs are posted every week, for example to avoid too many sticky jobs at one time. Sticky posts are limited to max 2 per company at one time. Job posts are a) published for 30 days, b) cannot be refunded, and c) renew automatically after 30 days unless you 1) disable auto-renew after posting on the edit page, or 2) close your job post on the edit page. We send a reminder 7 days by email before renewing. Renewing is the same price as the original job post for 30 days. If you buy a bundle, the discounted single job post price is used to renew. Automatic renewals can be self-refunded within 7 days after renewing with the link in the email.",
}
