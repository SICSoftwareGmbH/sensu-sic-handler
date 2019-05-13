// Copyright © 2019 SIC! Software GmbH

package output

import (
	sensu "github.com/sensu/sensu-go/types"
)

// Config configuration for handlers
type Config struct {
	SMTPAddress     string
	MailFrom        string
	SlackWebhookURL string
	SlackUsername   string
	SlackIconURL    string
	XMPPServer      string
	XMPPUsername    string
	XMPPPassword    string
}

// ExtendedEvent is a helper type for template resolution
type ExtendedEvent struct {
	Event            *sensu.Event
	Status           string
	EventAction      string
	EventKey         string
	Output           string
	FullOutput       string
	FormattedMessage string
}
