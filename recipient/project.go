// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/go-redis/redis"

	"sensu-sic-handler/redmine"
)

// ParseProject parse redmine project recipients
func ParseProject(redisClient *redis.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	switch len(args) {
	case 2:
		mails := readProjectMailsFromRedis(redisClient, args[0], args[1])

		for _, m := range mails {
			recipients = append(recipients, &Recipient{
				Type: HandlerTypeMail,
				ID:   fmt.Sprintf("mail|%s", m),
				Args: map[string]string{"mail": m},
			})
		}
	}

	return recipients
}

func readProjectMailsFromRedis(client *redis.Client, projectIdentifier string, roleIDStr string) []string {
	var mails []string

	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		return mails
	}

	val, err := client.Get(redmine.RedisKey(projectIdentifier, roleID, "mail")).Result()
	if err != nil {
		return mails
	}

	err = json.Unmarshal([]byte(val), &mails)
	if err != nil {
		return mails
	}

	return mails
}
