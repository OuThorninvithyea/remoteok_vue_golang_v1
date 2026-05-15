package store

import "backend/model/postJob"

var PostjobPreviewData = postJob.PreviewSection{
	Legend:     "PREVIEW",
	Heading:    "Here's a preview of how your job post will look like",
	Subheading: "Don't worry if it's not perfect the first time: your job is fully editable for free after posting it!",
	Card: postJob.PreviewCard{
		Logo:                   "/images/logo.png",
		CompanyPlaceholder:     "company",
		PositionPlaceholder:    "position",
		LocationBadge:          "🌏 Worldwide",
		HiringText:             "Company is hiring a",
		RoleText:               "Remote Position",
		DescriptionPlaceholder: `The description of the job position will appear here. Write this in the "Job Description" box above.`,
		ApplyHeading:           "How Do you apply?",
		ApplyPlaceholder:       `Here the instructions go on how to apply for this job. Write them in the "How to Apply?" box.`,
		ApplyButtonText:        "Apply for this job",
	},
	Watermark: "PREVIEW",
}
