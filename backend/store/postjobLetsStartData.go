package store

import "backend/model/postJob"

var PostjobLetsStartData = postJob.LetsStart{
	Legend: "LET'S START",
	Position: postJob.FormField{
		Label:       "POSITION*",
		Placeholder: "Position",
		Required:    true,
		HelpText:    `Please specify as single job position like "Marketing Manager" or "Node.js Developer", not a sentence like "Looking for PM / Biz Dev / Manager". We know your job is important but please DO NOT WRITE IN FULL CAPS. If posting multiple roles, please create multiple job posts. A job post is limited to a single job. We only allow real jobs, absolutely no MLM-type courses "learn how to work online" please.`,
	},
	CompanyName: postJob.FormField{
		Label:       "COMPANY NAME*",
		Placeholder: "Company name",
		Required:    true,
		HelpText:    "Your company's brand/trade name: without Inc., Ltd., B.V., Pte., etc.",
	},
	JobDescription: postJob.JobDescriptionField{
		Label:        "JOB DESCRIPTION*",
		AIButtonText: "✨ Write with AI",
	},
	EmploymentTypes: []postJob.SelectOption{
		{Value: "full_time", Label: "Full-time"},
		{Value: "part_time", Label: "Part-time"},
		{Value: "contractor", Label: "Contractor"},
		{Value: "temporary", Label: "Temporary"},
		{Value: "intern", Label: "Internship"},
		{Value: "per_diem", Label: "Per diem"},
		{Value: "volunteer", Label: "Volunteer"},
		{Value: "onsite", Label: "Onsite"},
	},
	PrimaryTags: []postJob.SelectOption{
		{Value: "", Label: "Select a primary tag", Disabled: true},
		{Value: "software_development", Label: "Software Development"},
		{Value: "customer_support", Label: "Customer Support"},
		{Value: "sales", Label: "Sales"},
		{Value: "marketing", Label: "Marketing"},
		{Value: "design", Label: "Design"},
		{Value: "front_end", Label: "Front End"},
		{Value: "back_end", Label: "Back End"},
		{Value: "legal", Label: "Legal"},
		{Value: "testing", Label: "Testing"},
		{Value: "quality_assurance", Label: "Quality Assurance"},
		{Value: "non_tech", Label: "Non-Tech"},
		{Value: "other", Label: "Other"},
	},
	PrimaryTagHelpText: "This primary tag shows first and increases visibility in the main sections. Your job is shown on every page that is tagged with it though. E.g. if you tag it as PHP, it shows for Remote PHP Jobs etc.",
	Tags: postJob.TagsField{
		Label:       "TAGS, KEYWORDS OR STACK*",
		Placeholder: "Type a tag or keyword to search and add it",
		HelpText:    "Short tags are preferred. Use tags like industry and tech stack. The first 3 or 4 tags are shown on the site, the other tags aren't but the job will be shown on each tag specific page (like /remote-react-jobs). We also sometimes generate tags automatically after you post/edit to supplement.",
	},
	Location: postJob.FormFieldWithNew{
		Label:         "JOB IS RESTRICTED TO LOCATIONS?",
		IsNew:         true,
		Placeholder:   "Type a location this job is restricted to like WorldWide, Europe, or Netherlands",
		Required:      true,
		HelpText:      `If you'd only like to hire people from a specific location or timezone this remote job is restricted to (e.g. Europe, United States or Japan). If not restricted, please leave it as "Worldwide". The less restricted this is, the more applicants you will get. Keeping it "Worldwide" is highly recommended as you'll have access to a worldwide pool of talent.`,
		HighlightText: "To promote fairness in remote work positions,",
		UnderlineText: "worldwide jobs are ranked higher.",
	},
}
