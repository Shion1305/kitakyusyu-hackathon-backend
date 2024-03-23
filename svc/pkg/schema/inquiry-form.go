package schema

import "errors"

type InquiryData struct {
	Firstname       string                 `json:"firstname"`
	Lastname        string                 `json:"lastname"`
	CompanyName     string                 `json:"company"`
	EmailAddress    string                 `json:"email"`
	BusinessContent string                 `json:"business"`
	InquiryDetails  string                 `json:"inquiry"`
	UseSlack        bool                   `json:"use_slack"`
	SlackInfo       *[]OtherSlackPersonnel `json:"slack_other"`
}

type OtherSlackPersonnel struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func (i InquiryData) Validate() error {
	if i.Firstname == "" {
		return errors.New("firstname is required")
	}
	if i.Lastname == "" {
		return errors.New("lastname is required")
	}
	if i.CompanyName == "" {
		return errors.New("company name is required")
	}
	if i.EmailAddress == "" {
		return errors.New("email address is required")
	}
	if i.UseSlack && (i.SlackInfo == nil || len(*i.SlackInfo) == 0) {
		return errors.New("slack info is required")
	}
	return nil
}
