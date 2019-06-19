// Copyright © 2019 SIC! Software GmbH
// Adapted from https://github.com/sensu/sensu-slack-handler

package output

import (
	"errors"

	"github.com/bluele/slack"

	"github.com/SICSoftwareGmbH/sensu-sic-handler/recipient"
)

// Slack handles slack recipients (recipient.HandlerTypeSlack)
func Slack(recipient *recipient.Recipient, event *ExtendedEvent, config *Config) error {
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

func slackMessageColor(event *ExtendedEvent) string {
	switch event.Event.Check.Status {
	case 0:
		return "good"
	case 2:
		return "danger"
	default:
		return "warning"
	}
}

func slackMessageAttachment(event *ExtendedEvent) *slack.Attachment {
	return &slack.Attachment{
		Title:    "Description",
		Text:     event.Event.Check.Output,
		Fallback: event.FormattedMessage,
		Color:    slackMessageColor(event),
		Fields: []*slack.AttachmentField{
			{
				Title: "Status",
				Value: messageEventStatus(event.Event),
				Short: false,
			},
			{
				Title: "Entity",
				Value: event.Event.Entity.Name,
				Short: true,
			},
			{
				Title: "Check",
				Value: event.Event.Check.Name,
				Short: true,
			},
		},
	}
}
