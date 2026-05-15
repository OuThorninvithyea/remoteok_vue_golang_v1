package postJob

type TrustedLogo struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type TrustedCompanies struct {
	HeaderText    string        `json:"headerText"`
	HighlightText string        `json:"highlightText"`
	Logos         []TrustedLogo `json:"logos"`
	ShowMoreText  string        `json:"showMoreText"`
}

type SidebarStat struct {
	Label         string `json:"label"`
	Value         string `json:"value"`
	Sub           string `json:"sub"`
	Type          string `json:"type"`
	StarIcon      string `json:"starIcon"`
	StarCount     int    `json:"starCount"`
	Image         string `json:"image"`
	ImageAlt      string `json:"imageAlt"`
	ApiLink       string `json:"apiLink"`
	HasGoogleIcon bool   `json:"hasGoogleIcon"`
}

type Testimonial struct {
	Image           string `json:"image"`
	Alt             string `json:"alt"`
	BeforeHighlight string `json:"beforeHighlight"`
	Highlight       string `json:"highlight"`
	Muted           string `json:"muted"`
	Name            string `json:"name"`
	Role            string `json:"role"`
	Company         string `json:"company"`
}

type Founder struct {
	Image   string   `json:"image"`
	Alt     string   `json:"alt"`
	Bio     []string `json:"bio"`
	Name    string   `json:"name"`
	Role    string   `json:"role"`
	Company string   `json:"company"`
}

type PostJobSidebar struct {
	TrustedCompanies TrustedCompanies `json:"trustedCompanies"`
	Stats            []SidebarStat    `json:"stats"`
	Testimonials     []Testimonial    `json:"testimonials"`
	Founder          Founder          `json:"founder"`
}
