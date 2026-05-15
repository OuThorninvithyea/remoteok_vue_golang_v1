package store

import "backend/model/postJob"

var PostjobSidebarData = postJob.PostJobSidebar{
	TrustedCompanies: postJob.TrustedCompanies{
		HeaderText:    "Remote OK is the most popular remote jobs board in the world trusted by millions of remote workers and leading remote companies like",
		HighlightText: "the most popular remote jobs board in the world",
		Logos: []postJob.TrustedLogo{
			{Src: "/images/aws.png", Alt: "AWS"},
			{Src: "/images/espn.svg", Alt: "ESPN"},
			{Src: "/images/ibm.png", Alt: "IBM"},
			{Src: "/images/microsoft.png", Alt: "Microsoft"},
			{Src: "/images/scale-ai.png", Alt: "Scale AI"},
			{Src: "/images/ycombinator.png", Alt: "Y Combinator"},
			{Src: "/images/github.png", Alt: "GitHub"},
			{Src: "/images/posthog.png", Alt: "PostHog"},
			{Src: "/images/mozilla.png", Alt: "Mozilla"},
			{Src: "/images/stripe.png", Alt: "Stripe"},
			{Src: "/images/cloudflare.png", Alt: "Cloudflare"},
			{Src: "/images/shopify.png", Alt: "Shopify"},
			{Src: "/images/godaddy.png", Alt: "GoDaddy"},
			{Src: "/images/easyjet.png", Alt: "easyJet"},
			{Src: "/images/indeed.png", Alt: "Indeed"},
		},
		ShowMoreText: "SHOW MORE",
	},
	Stats: []postJob.SidebarStat{
		{Label: "Starting from", Value: "💳 $299", Sub: "for 30 days"},
		{Label: "Reach", Value: "🚀 2,700,000+", Sub: "remote workers/mo"},
		{Label: "Distributed on the", Value: "Google for Jobs", Sub: "recruitment network", HasGoogleIcon: true},
		{Label: "Rated", Type: "stars", StarIcon: "/images/emoji-star.png", StarCount: 5, Sub: "Customer rating <strong>9.0</strong> | <strong>60,259</strong> reviews"},
		{Label: "Emailed to", Value: "📬 2,309,564", Sub: "remote job seekers"},
		{Label: "Crossposted to", Value: "✨ 240 job boards", Sub: "that currently use our API", ApiLink: "#"},
		{Label: "Guaranteed", Value: "🎡 200+ [Apply] clicks", Sub: "or we auto bump it for free"},
		{Label: "Pay with", Type: "image", Image: "/images/payment-methods.png", ImageAlt: "Payment methods", Sub: "🔒 Secure payment with Stripe"},
	},
	Testimonials: []postJob.Testimonial{
		{Image: "/images/testimonial-edwin-stripe-3.jpeg", Alt: "Edwin", Highlight: "The response has been amazing! A lotttt of applicants.", Muted: `Thank you, and everything Remote OK did, to help with this!"`, Name: "Edwin", Role: "DevRel", Company: "OpenAI"},
		{Image: "/images/testimonial-sara-komoot-2.jpg", Alt: "Sara", Highlight: "Remote OK has been an essential platform for attracting great talent to our remote-first company", Name: "Sara", Role: "Recruitment", Company: "Komoot"},
		{Image: "/images/testimonial-andy-tettra.jpg", Alt: "Andy", Highlight: "I want to say that having 1,500 applicants to sort through is a good problem to have", Muted: `and thanks to y'all for making that happen"`, Name: "Andy", Role: "CEO", Company: "Tettra"},
		{Image: "/images/testimonial-tris-aula.jpg", Alt: "Tris", Highlight: "We are loving the performance of the job ads, we have had so many great applicants.", Muted: `Thank you for your great work on building this job board :)"`, Name: "Tris", Role: "Global Sourcing Lead", Company: "Aula Education"},
		{Image: "/images/testimonial-zsolt.jpeg", Alt: "Zsolt", Highlight: "I love the site, I tried so many different job sites and Remote OK is by far the best to deliver", Name: "Zsolt", Role: "Founder & CEO", Company: "PingPong"},
		{Image: "/images/testimonial-baptiste-crisp.jpg", Alt: "Baptiste", Highlight: "We super like Remote OK. We had around 100 applicants for our previous job post,", Muted: `so we are doing a new one :)"`, Name: "Baptiste", Role: "Founder & CEO", Company: "Crisp Chat"},
		{Image: "/images/testimonial-ken-savvy.jpg", Alt: "Ken", Highlight: "Awesome, you rock! Your customer support is a perfect example of why I stick with your site", Muted: `as our "go-to" job board..."`, Name: "Ken", Role: "CEO", Company: "Savvy"},
		{Image: "/images/testimonial-peter-benei.png", Alt: "Peter", BeforeHighlight: "We posted 2 jobs in the recent month now, and a quick feedback on the applicants: we're not just getting a ton of them, but", Highlight: "people coming from Remote OK are much much much more valid and higher quality", Muted: `compared to people coming from LinkedIn, Angel, VirtualVocations, Emotive, or Indeed, as we posted our jobs there too. So big fan of Remote OK, the last job was only posted there. Thanks for the good stuff!"`, Name: "Peter", Role: "COO", Company: "KISSPatent"},
		{Image: "/images/testimonial-crossover.jpg", Alt: "Alex", Highlight: "It was the best job board targeted at remote professionals in terms of results.", Muted: `Thank you for the collaboration, it was a pleasure working with you and with Remote OK."`, Name: "Alex", Role: "Global Sourcing Strategist", Company: "Crossover"},
	},
	Founder: postJob.Founder{
		Image: "/images/img_1.png",
		Alt:   "Pieter Levels",
		Bio: []string{
			`👋 Hi! I'm the maker of Remote OK and other sites related to remote work such as <a href="https://nomads.com">Nomads.com</a>. Remote OK isn't a big team, it's actually just a one-man operation which is me on a laptop somewhere in the world. I built Remote OK to help 🚀 accelerate the revolution that is remote work (and pay my rent and coffee). My site has now become the <a href="https://remoteok.com">#1 remote job platform</a> in the world!`,
			"Remote work gives people more flexibility in their daily lives, and lets employers hire the 🏆 the best talent regardless of where they are located in the 🌍 world. That's the most significant change to work since the industrial revolution, and that's why I make this site.",
			`Need help posting your job? Here's <a href="https://x.com/levelsio">my 𝕏</a>. Tweet to me, and I'll help you personally and can also help you if you'd like to buy multiple job posts packages or get a discount. Go remote!`,
		},
		Name:    "Pieter Levels",
		Role:    "Founder",
		Company: "Remote OK",
	},
}
