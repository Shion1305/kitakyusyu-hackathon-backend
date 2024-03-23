package uc

import (
	"errors"
	"fmt"
	"kitakyusyu-hackathon/pkg/slack"
)

type InviteSlack struct {
	slack slack.Slack
}

func NewInviteSlack(slack slack.Slack) *InviteSlack {
	return &InviteSlack{slack: slack}
}

type InviteSlackInput struct {
	ChannelName string
	StaffIDs    []string
	GuestInfo   []GuestInfo
}

type GuestInfo struct {
	Firstname string
	Lastname  string
	Email     string
}

type InviteSlackOutput struct {
	ChannelName string
	ChannelLink string
}

func (uc InviteSlack) Do(input InviteSlackInput) (*InviteSlackOutput, error) {
	if len(input.StaffIDs) == 0 || len(input.GuestInfo) == 0 {
		return nil, errors.New("no staff or guest info provided")
	}

	// first create conversation (channel)
	channel, err := uc.slack.CreateChannel(slack.CreateConversationParams{
		ChannelName: input.ChannelName,
		IsPrivate:   true,
	})
	if err != nil {
		return nil, err
	}

	// then issue invitation link
	if err := uc.slack.InviteToChannel(channel.ID, input.StaffIDs); err != nil {
		return nil, err
	}

	errInfo := make(map[string]error)
	for _, g := range input.GuestInfo {
		// invite guest
		if err := uc.slack.InviteGuestToConversation(channel.ID, g.Firstname, g.Lastname, g.Email); err != nil {
			errInfo[g.Email] = err
		}
	}
	if len(errInfo) > 0 {
		errMsg := "failed to invite guests: "
		for email, err := range errInfo {
			errMsg += email + ":" + err.Error() + ", "
		}
		return &InviteSlackOutput{
			ChannelName: input.ChannelName,
			ChannelLink: genConversationLink(channel.ID),
		}, errors.New(errMsg)
	}
	return &InviteSlackOutput{
		ChannelName: input.ChannelName,
		ChannelLink: genConversationLink(channel.ID),
	}, nil
}

func genConversationLink(channelID string) string {
	return fmt.Sprintf("https://slack.com/app_redirect?channel=%s", channelID)
}
