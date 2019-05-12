// Copyright © 2019 SIC! Software GmbH

package handler

import (
	"os"
	"fmt"

	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

// Handle handles recipients
func Handle(recipients []*recipient.Recipient, event *sensu.Event, config *Config) error {
	recipientMap := make(map[string]bool)

	for _, rcpt := range recipients {
		if _, ok := recipientMap[rcpt.ID]; !ok {
			var err error
			err = nil

			switch rcpt.Type {
			case recipient.HandlerTypeNone:
			case recipient.HandlerTypeMail:
				err = HandleMail(rcpt, event, config)
			case recipient.HandlerTypeXMPP:
				err = HandleXMPP(rcpt, event, config)
			case recipient.HandlerTypeSlack:
				err = HandleSlack(rcpt, event, config)
			default:
				fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupported handler: %s", rcpt))
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("failed to handle: %s", err.Error()))
			}

			recipientMap[rcpt.ID] = true
		}
	}

	return nil
}
