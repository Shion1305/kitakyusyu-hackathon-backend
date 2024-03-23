package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kitakyusyu-hackathon/pkg/config"
	"kitakyusyu-hackathon/pkg/sendgrid"
	"kitakyusyu-hackathon/pkg/slack"
	"kitakyusyu-hackathon/svc/pkg/schema"
	"kitakyusyu-hackathon/svc/pkg/uc"
	"log"
)

type InquiryHandler struct {
	slackClient *slack.Slack
	inviteUC    *uc.InviteSlack
	sendgrid    *sendgrid.Sendgrid
}

func NewInquiryHandler() *InquiryHandler {
	s := slack.NewSlack()
	conf := config.Get()
	return &InquiryHandler{
		slackClient: &s,
		inviteUC:    uc.NewInviteSlack(s, conf.Slack.TeamID),
		sendgrid:    sendgrid.NewSendgrid(),
	}
}

func (h *InquiryHandler) HandleInquiry() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("inquiry request")
		data := schema.InquiryData{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			log.Printf("failed to bind json, err: %v\n", err)
			return
		}

		if err := data.Validate(); err != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			log.Printf("validation error: %v\n", err)
			return
		}

		log.Printf("inquiry data: %+v\n", data)

		if data.UseSlack {
			h.handleSlack(data)
		} else {
			h.handleMail(data)
		}

		log.Printf("inquery process succeeded\n")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}

func (h InquiryHandler) handleSlack(data schema.InquiryData) {
	var guests []uc.GuestInfo
	if *data.SlackInfo != nil {
		guests = make([]uc.GuestInfo, 0, len(*data.SlackInfo)+1)
		for _, s := range *data.SlackInfo {
			guests = append(guests, uc.GuestInfo{
				Email:     s.Email,
				Firstname: s.Firstname,
				Lastname:  s.Lastname,
			})
		}
	}
	guests = append(guests, uc.GuestInfo{
		Email:     data.EmailAddress,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
	})
	inviteInput := uc.InviteSlackInput{
		ChannelName: data.CompanyName,
		StaffIDs:    []string{"U04936U1UEB"},
		GuestInfo:   guests,
	}
	inviteResult, err := h.inviteUC.Do(inviteInput)
	if err != nil {
		log.Printf("failed to invite slack, err: %v\n", err)
		return
	}
	log.Printf("channel created: %s\n", inviteResult.ChannelName)
	log.Printf("channel link: %s\n", inviteResult.ChannelLink)
}

func (h InquiryHandler) handleMail(data schema.InquiryData) {
	h.sendgrid.SendMailNotify(fmt.Sprintf("%s %s", data.Firstname, data.Lastname), data.EmailAddress)
}
