// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// ParseXMPP parse mail recipients (HandlerTypeXMPP)
func ParseXMPP(redisClient *redis.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	if len(args) == 2 {
		switch args[0] {
		case "muc":
			recipients = append(recipients, &Recipient{
				Type: HandlerTypeXMPP,
				ID:   fmt.Sprintf("xmpp|muc|%s", args[1]),
				Args: map[string]string{"type": "muc", "room": args[1]},
			})
		case "user":
			recipients = append(recipients, &Recipient{
				Type: HandlerTypeXMPP,
				ID:   fmt.Sprintf("xmpp|user|%s", args[1]),
				Args: map[string]string{"type": "user", "user": args[1]},
			})
		}
	}

	return recipients
}
