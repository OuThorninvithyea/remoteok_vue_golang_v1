package store

import "backend/model/login"

var LoginData = login.LoginData{
	BackgroundImg:    "/images/promotion.jpg",
	Heading:          `Find a remote job on<br>the <a href="#" class="highlight-text text-underline"> #1 remote work platform</a>`,
	InputPlaceholder: "Username or email",
	EmailButton:      "Continue with email",
	DividerText:      "or",
	GoogleButton:     "Continue with Google",
	SignupText:        "Not a members yet?",
	SignupLink:        "Sign up",
	TermsText:        "By logging in you agree to our",
	TermsLink:        "term and services",
}
