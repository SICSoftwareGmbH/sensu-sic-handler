// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	"fmt"
)

// EtcdKey generate key for etcd project lookups
func EtcdKey(projectIdentifier string, roleID int, suffix string) string {
	return fmt.Sprintf("/sensu.sic.software/project/%s/%d/%s", projectIdentifier, roleID, suffix)
}
