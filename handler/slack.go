// Copyright © 2019 SIC! Software GmbH
// Adapted from https://github.com/sensu/sensu-slack-handler

package handler

import (
	"errors"

	sensu "github.com/sensu/sensu-go/types"
	"github.com/bluele/slack"

	"sensu-sic-handler/recipient"
)

// HandleSlack handles slack recipients (recipient.HandlerTypeSlack)
func HandleSlack(recipient *recipient.Recipient, event *sensu.Event, config *Config) error {
	if len(config.SlackWebhookURL) == 0 {
		return errors.New("webhook url is empty")
	}

	hook := slack.NewWebHook(config.SlackWebhookURL)

	return hook.PostMessage(&slack.WebHookPostPayload{
		Attachments: []*slack.Attachment{slackMessageAttachment(event)},
		Channel:     recipient.Args["channel"],
		IconUrl:     config.SlackIconURL,
		Username:    config.SlackUsername,
	})
}

func slackMessageColor(event *sensu.Event) string {
	switch event.Check.Status {
	case 0:
		return "good"
	case 2:
		return "danger"
	default:
		return "warning"
	}
}

func slackMessageAttachment(event *sensu.Event) *slack.Attachment {
	return &slack.Attachment{
		Title:    "Description",
		Text:     event.Check.Output,
		Fallback: formattedMessage(event),
		Color:    slackMessageColor(event),
		Fields: []*slack.AttachmentField{
			&slack.AttachmentField{
				Title: "Status",
				Value: messageStatus(event),
				Short: false,
			},
			&slack.AttachmentField{
				Title: "Entity",
				Value: event.Entity.Name,
				Short: true,
			},
			&slack.AttachmentField{
				Title: "Check",
				Value: event.Check.Name,
				Short: true,
			},
		},
	}
}
