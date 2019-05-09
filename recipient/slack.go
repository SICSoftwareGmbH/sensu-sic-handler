// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// ParseSlack parse mail recipients (HandlerTypeSlack)
func ParseSlack(redisClient *redis.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	switch len(args) {
	case 2:
		switch args[0] {
		case "channel":
			recipients = append(recipients, &Recipient{
				Type: HandlerTypeSlack,
				ID:   fmt.Sprintf("slack|channel|%s", args[1]),
				Args: map[string]string{"type": "channel", "channel": args[1]},
			})
		}
	}

	return recipients
}
