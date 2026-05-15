package store

import "backend/model/postJob"

var PostjobFormData = postJob.PostJobForm{
	ToolbarIcons: PostjobToolbarData,
	LetsStart:    PostjobLetsStartData,
	DesignPost:   PostjobDesignPostData,
	JobDetails:   PostjobJobDetailsData,
	Company:      PostjobCompanyData,
	Feedback:     PostjobFeedbackData,
	Preview:      PostjobPreviewData,
	Partner:      PostjobPartnerData,
	Sidebar:      PostjobSidebarData,
	Footer:       PostjobFooterData,
}
