package postJob

type PostJobForm struct {
	ToolbarIcons []ToolbarIconGroup `json:"toolbarIcons"`
	LetsStart    LetsStart          `json:"letsStart"`
	DesignPost   DesignPost         `json:"designPost"`
	JobDetails   JobDetails         `json:"jobDetails"`
	Company      Company            `json:"company"`
	Feedback     FeedbackSection    `json:"feedback"`
	Preview      PreviewSection     `json:"preview"`
	Partner      PartnerSection     `json:"partner"`
	Sidebar      PostJobSidebar     `json:"sidebar"`
	Footer       PostJobFooter      `json:"footer"`
}
