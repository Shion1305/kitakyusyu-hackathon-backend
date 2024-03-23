package slack

import (
	"github.com/slack-go/slack"
	"kitakyusyu-hackathon/pkg/config"
)

type Slack struct {
	appID             string
	clientID          string
	clientSecret      string
	signingSecret     string
	verificationToken string
	userOauthToken    string
	teamID            string
	client            *slack.Client
}

func NewSlack() Slack {
	conf := config.Get()

	client := slack.New(conf.Slack.UserOAuthToken)
	return Slack{
		appID:             conf.Slack.AppID,
		clientID:          conf.Slack.ClientID,
		clientSecret:      conf.Slack.ClientSecret,
		signingSecret:     conf.Slack.SigningSecret,
		verificationToken: conf.Slack.VerificationToken,
		userOauthToken:    conf.Slack.UserOAuthToken,
		teamID:            conf.Slack.TeamID,
		client:            client,
	}
}

type CreateConversationParams struct {
	ChannelName string
	IsPrivate   bool
}

type CreateConversationResponse struct {
}

func (s Slack) CreateChannel(param CreateConversationParams) (*slack.Channel, error) {
	return s.client.CreateConversation(
		slack.CreateConversationParams{
			ChannelName: param.ChannelName,
			IsPrivate:   param.IsPrivate,
			TeamID:      s.teamID,
		})
}

func (s Slack) InviteToChannel(channelID string, userIDs []string) error {
	_, err := s.client.InviteUsersToConversation(channelID, userIDs...)
	return err
}

func (s Slack) CreateInviteLink(channelID string) (string, error) {
	return s.client.GetPermalink(&slack.PermalinkParameters{
		Channel: channelID,
	})
}

func (s Slack) InviteGuestToConversation(
	channelID, firstname, lastname, email string,
) error {
	return s.client.InviteGuest(s.teamID, channelID, firstname, lastname, email)
}
