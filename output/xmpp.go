// Copyright © 2019 SIC! Software GmbH

package output

import (
	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

// XMPP handles XMPP recipients (recipient.HandlerTypeXMPP)
func XMPP(recipient *recipient.Recipient, event *sensu.Event, config *Config) error {
	return nil
}
