package postJob

type PostJobLogoMark struct {
	Text      string `json:"text"`
	Badge     string `json:"badge"`
	Trademark string `json:"trademark"`
	Link      string `json:"link"`
}

type PostJobButton struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type PostJobHeader struct {
	Logo   PostJobLogoMark `json:"logo"`
	Title  string          `json:"title"`
	Button PostJobButton   `json:"button"`
}
