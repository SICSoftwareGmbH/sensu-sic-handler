// Copyright © 2019 SIC! Software GmbH

package output

import (
	"fmt"
	"os"

	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

// Notify handles recipients
func Notify(recipients []*recipient.Recipient, event *sensu.Event, config *Config) error {
	recipientMap := make(map[string]bool)

	extendedEvent := extendedEventFromEvent(event)

	for _, rcpt := range recipients {
		if _, ok := recipientMap[rcpt.ID]; !ok {
			var err error

			switch rcpt.Type {
			case recipient.OutputTypeNone:
			case recipient.OutputTypeMail:
				err = Mail(rcpt, extendedEvent, config)
			case recipient.OutputTypeXMPP:
				err = XMPP(rcpt, extendedEvent, config)
			case recipient.OutputTypeSlack:
				err = Slack(rcpt, extendedEvent, config)
			default:
				fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupported handler: %q", rcpt))
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("failed to handle: %s", err.Error()))
			}

			recipientMap[rcpt.ID] = true
		}
	}

	return nil
}
