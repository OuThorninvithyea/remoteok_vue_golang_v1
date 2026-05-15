package store

import "backend/model/postJob"

var PostjobCompanyData = postJob.Company{
	Legend: "COMPANY",
	Fields: []postJob.CompanyField{
		{Label: "COMPANY TWITTER", Type: "text", Placeholder: "e.g. @remoteok", Required: false, HelpText: "Twitter handle without @. Used to tag your company when we tweet your job post."},
		{Label: "COMPANY EMAIL* (PRIVATE, FOR INVOICES + EDIT LINK)", Type: "email", Placeholder: "Email", Required: true, HelpText: "This is where we send the receipt and edit link. It is NOT shown on the job post."},
		{Label: "INVOICE EMAIL (PRIVATE)", Type: "email", Placeholder: "Invoice Email", Required: false, HelpText: "If you'd like the invoice to be sent to a different email address (like your accounting department)."},
		{Label: "INVOICE ADDRESS", Type: "textarea", Placeholder: "e.g. your company's full legal name and full invoice address...", Required: false, HelpText: "This is shown on the invoice.", TextareaClass: "textarea-small"},
		{Label: "INVOICE NOTES / PO NUMBER", Type: "text", Placeholder: "e.g. PO 12345", Required: false, HelpText: "Optional: any notes you want to show on the invoice."},
	},
	PayLaterLabel: "Pay later",
	ShowLogoCheckbox: postJob.ShowLogoCheckbox{
		Name:    "show_company_logo",
		Text:    "Show my ⭐️ company logo besides my post (+$44)",
		Checked: true,
		Price:   44,
	},
}
