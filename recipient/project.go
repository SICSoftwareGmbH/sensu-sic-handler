// Copyright © 2019 SIC! Software GmbH

package recipient

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	etcd "go.etcd.io/etcd/clientv3"

	"sensu-sic-handler/redmine"
)

// ParseProject parse redmine project recipients
func ParseProject(etcdClient *etcd.Client, value string) []*Recipient {
	recipients := make([]*Recipient, 0)

	args := strings.Split(value, ":")

	if len(args) == 2 {
		mails := readProjectMailsFromEtcd(etcdClient, args[0], args[1])

		for _, m := range mails {
			recipients = append(recipients, &Recipient{
				Type: OutputTypeMail,
				ID:   fmt.Sprintf("mail|%s", m),
				Args: map[string]string{"mail": m},
			})
		}
	}

	return recipients
}

func readProjectMailsFromEtcd(client *etcd.Client, projectIdentifier string, roleIDStr string) []string {
	var mails []string

	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		return mails
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := client.Get(ctx, redmine.EtcdKey(projectIdentifier, roleID, "mail"))
	cancel()
	if err != nil {
		return mails
	}

	if len(resp.Kvs) != 1 {
		return mails
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &mails)
	if err != nil {
		return mails
	}

	return mails
}
