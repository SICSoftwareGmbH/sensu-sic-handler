// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"os"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// Parse parses recipients from string
func Parse(redisClient *redis.Client, value string) []*Recipient {
	recipientDefinitions := strings.Split(value, ",")
	recipients := make([]*Recipient, 0)

	for _, recipientDefinition := range recipientDefinitions {
		val := strings.SplitN(recipientDefinition, ":", 2)

		switch val[0] {
		case "slack":
			recipients = append(recipients, ParseSlack(redisClient, val[1])...)
		case "xmpp":
			recipients = append(recipients, ParseXMPP(redisClient, val[1])...)
		case "mail":
			recipients = append(recipients, ParseMail(redisClient, val[1])...)
		case "project":
			recipients = append(recipients, ParseProject(redisClient, val[1])...)
		default:
			fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupported recipient: %s", recipientDefinition))
		}
	}

	return recipients
}
