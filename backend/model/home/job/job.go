package job

type JobDetail struct {
	HiringLabel      string   `json:"hiringLabel"`
	RoleTitle        string   `json:"roleTitle"`
	Paragraphs       []string `json:"paragraphs"`
	Responsibilities []string `json:"responsibilities"`
	Offers           []string `json:"offers"`
	InterviewProcess []string `json:"interviewProcess"`
	InterviewNotice  string   `json:"interviewNotice"`
	Salary           string   `json:"salary"`
	ApplyQuestion    string   `json:"applyQuestion"`
	ApplyFormUrl     string   `json:"applyFormUrl"`
	ApplyFormLabel   string   `json:"applyFormLabel"`
	Text             string   `json:"text"`
	Benefits         []string `json:"benefits"`
	CompanyName      string   `json:"companyName"`
	CompanyWebsite   string   `json:"companyWebsite"`
	CompanyLogo      string   `json:"companyLogo"`
	QrCode           string   `json:"qrCode"`
	QrLink           string   `json:"qrLink"`
	ApplyUrl         string   `json:"applyUrl"`
}

type Job struct {
	BadgeClass    string     `json:"badgeClass"`
	Type          string     `json:"type"`
	CardClass     string     `json:"cardClass"`
	Logo          string     `json:"logo"`
	Title         string     `json:"title"`
	Company       string     `json:"company"`
	IsNew         bool       `json:"isNew"`
	LocationBadge string     `json:"locationBadge"`
	SalaryBadge   string     `json:"salaryBadge"`
	TypeBadge     string     `json:"typeBadge"`
	MiddleBadges  []string   `json:"middleBadges"`
	ExtraBadges   []string   `json:"extraBadges"`
	Time          string     `json:"time"`
	IsPinned      bool       `json:"isPinned"`
	ApplyLink     string     `json:"applyLink"`
	ApplyLabel    string     `json:"applyLabel"`
	Img           string     `json:"img"`
	Paragraph     string     `json:"paragraph"`
	Rating        string     `json:"rating"`
	Button        string     `json:"button"`
	Detail        *JobDetail `json:"detail"`
}
