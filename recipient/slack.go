// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// ParseSlack parse mail recipients (OutputTypeSlack)
func ParseSlack(redisClient *redis.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	if len(args) == 2 {
		if args[0] == "channel" {
			recipients = append(recipients, &Recipient{
				Type: OutputTypeSlack,
				ID:   fmt.Sprintf("slack|channel|%s", args[1]),
				Args: map[string]string{"type": "channel", "channel": args[1]},
			})
		}
	}

	return recipients
}
