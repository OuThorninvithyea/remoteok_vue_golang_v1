package login

type LoginData struct {
	BackgroundImg    string `json:"backgroundImg"`
	Heading          string `json:"heading"`
	InputPlaceholder string `json:"inputPlaceholder"`
	EmailButton      string `json:"emailButton"`
	DividerText      string `json:"dividerText"`
	GoogleButton     string `json:"googleButton"`
	SignupText        string `json:"signupText"`
	SignupLink        string `json:"signupLink"`
	TermsText        string `json:"termsText"`
	TermsLink        string `json:"termsLink"`
}
