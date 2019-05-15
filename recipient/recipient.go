// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"fmt"
	"os"
	"strings"

	etcd "go.etcd.io/etcd/clientv3"
)

// Parse parses recipients from string
func Parse(etcdClient *etcd.Client, value string) []*Recipient {
	recipientDefinitions := strings.Split(value, ",")
	recipients := make([]*Recipient, 0)

	for _, recipientDefinition := range recipientDefinitions {
		val := strings.SplitN(recipientDefinition, ":", 2)

		switch val[0] {
		case "slack":
			recipients = append(recipients, ParseSlack(etcdClient, val[1])...)
		case "xmpp":
			recipients = append(recipients, ParseXMPP(etcdClient, val[1])...)
		case "mail":
			recipients = append(recipients, ParseMail(etcdClient, val[1])...)
		case "project":
			recipients = append(recipients, ParseProject(etcdClient, val[1])...)
		default:
			fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupported recipient: %s", recipientDefinition))
		}
	}

	return recipients
}
