package uc

import (
	"errors"
	"fmt"
	"kitakyusyu-hackathon/pkg/slack"
	"log"
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
	if len(input.GuestInfo) == 0 {
		return nil, errors.New("no guest info provided")
	}

	// first create conversation (channel)
	channel, err := uc.slack.CreateChannel(slack.CreateConversationParams{
		ChannelName: input.ChannelName,
		IsPrivate:   true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}

	if len(input.StaffIDs) > 0 {
		if err := uc.slack.InviteToChannel(channel.ID, input.StaffIDs); err != nil {
			return nil, fmt.Errorf("failed to invite staff to channel: %w", err)
		}
	}

	guests := make([]string, 0, len(input.GuestInfo))
	for _, guest := range input.GuestInfo {
		guests = append(guests, guest.Email)
	}
	conversation, b, err := uc.slack.InviteGuestToConversation(channel.ID, guests)
	if err != nil {
		log.Printf("failed to invite guest to channel: %v, %v, %v", conversation, b, err)
	}

	return &InviteSlackOutput{
		ChannelName: input.ChannelName,
		ChannelLink: genConversationLink(channel.ID),
	}, nil
}

func genConversationLink(channelID string) string {
	return fmt.Sprintf("https://slack.com/app_redirect?channel=%s", channelID)
}
