// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	"fmt"
)

// RedisKey generate key for redis project lookups
func RedisKey(projectIdentifier string, roleID int, suffix string) string {
	return fmt.Sprintf("project.%s.%d.%s", projectIdentifier, roleID, suffix)
}
