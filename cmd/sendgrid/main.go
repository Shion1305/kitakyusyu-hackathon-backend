package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"kitakyusyu-hackathon/pkg/config"
	"log"
)

func main() {
	conf := config.Get()
	from := mail.NewEmail("Example User", "shion1305@proton.me")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example User", "shion1305@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(conf.SendGrid.APIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
