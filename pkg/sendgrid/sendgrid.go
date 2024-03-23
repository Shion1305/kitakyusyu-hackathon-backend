package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"kitakyusyu-hackathon/pkg/config"
	"log"
)

type Sendgrid struct {
	api_key string
}

func NewSendgrid() *Sendgrid {
	conf := config.Get()
	return &Sendgrid{
		api_key: conf.SendGrid.APIKey,
	}
}

func (s *Sendgrid) SendMail() {

}

func (s *Sendgrid) SendMailNotify(name, to string) {
	log.Printf("sending mail to %s\n", to)
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("JCHキャンペーン運営事務局", "shion1305@proton.me"))
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(name+" 様", to))
	m.SetTemplateID("d-9ea807833b634a778084e55411b0cbcf")
	m.AddPersonalizations(p)

	req := sendgrid.GetRequest(s.api_key, "/v3/mail/send", "https://api.sendgrid.com")
	req.Method = "POST"
	req.Body = mail.GetRequestBody(m)
	resp, err := sendgrid.API(req)
	if err != nil {
		log.Printf("failed to send mail, err: %v\n", err)
	}
	log.Printf("response: %v\n", resp)
}

func (s *Sendgrid) SendMailSlack(name, to string) {
	log.Printf("sending mail to %s\n", to)
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("JCHキャンペーン運営事務局", "shion1305@proton.me"))
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(name+" 様", to))
	m.SetTemplateID("d-9ea807833b634a778084e55411b0cbcf")
	m.AddPersonalizations(p)

	req := sendgrid.GetRequest(s.api_key, "/v3/mail/send", "https://api.sendgrid.com")
	req.Method = "POST"
	req.Body = mail.GetRequestBody(m)
	resp, err := sendgrid.API(req)
	if err != nil {
		log.Printf("failed to send mail, err: %v\n", err)
	}
	log.Printf("response: %v\n", resp)
}
