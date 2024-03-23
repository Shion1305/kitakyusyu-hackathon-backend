package schema

type InquiryData struct {
	Name           string `json:"name"`
	CompanyName    string `json:"company"`
	EmailAddress   string `json:"email"`
	InquiryDetails string `json:"inquiry"`
}
