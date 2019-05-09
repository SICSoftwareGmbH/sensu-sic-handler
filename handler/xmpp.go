// Copyright © 2019 SIC! Software GmbH

package handler

import (
	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

// HandleXMPP handles XMPP recipients (recipient.HandlerTypeXMPP)
func HandleXMPP(recipient *recipient.Recipient, event *sensu.Event, config *Config) error {
	return nil
}
