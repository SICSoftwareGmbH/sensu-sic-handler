// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"strings"

	etcd "go.etcd.io/etcd/clientv3"
)

// ParseSlack parse mail recipients (OutputTypeSlack)
func ParseSlack(etcdClient *etcd.Client, value string) []*Recipient {
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
