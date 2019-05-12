// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

// ParseMail parse mail recipients (OutputTypeMail)
func ParseMail(redisClient *redis.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	if len(args) == 1 {
		recipients = append(recipients, &Recipient{
			Type: OutputTypeMail,
			ID:   fmt.Sprintf("mail|%s", args[0]),
			Args: map[string]string{"mail": args[0]},
		})
	}

	return recipients
}
