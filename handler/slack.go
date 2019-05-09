// Copyright © 2019 SIC! Software GmbH

package handler

import (
	sensu "github.com/sensu/sensu-go/types"

	"sensu-sic-handler/recipient"
)

// HandleSlack handles slack recipients (recipient.HandlerTypeSlack)
func HandleSlack(recipient *recipient.Recipient, event *sensu.Event, config *Config) error {
	return nil
}
