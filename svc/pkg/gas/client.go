package gas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kitakyusyu-hackathon/pkg/config"
	"net/http"
)

type InquiryData struct {
	Firstname       string    `json:"firstname"`
	Lastname        string    `json:"lastname"`
	CompanyName     string    `json:"company"`
	EmailAddress    string    `json:"email"`
	Purpose         string    `json:"purpose"`
	InquiryDetails  string    `json:"inquiry"`
	UseSlack        bool      `json:"use_slack"`
	SlackChannelURL string    `json:"slack_channel_url"`
	SlackInfoEmails *[]string `json:"slack_other"`
}
type GAS struct {
	appUrl string
}

func NewGAS() GAS {
	conf := config.Get()
	return GAS{
		appUrl: conf.GAS.AppURL,
	}
}
func (g GAS) PostData(data InquiryData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	resp, err := http.Post(g.appUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error sending request to GAS: %s\n", err)
		return
	}
	defer resp.Body.Close()
}
