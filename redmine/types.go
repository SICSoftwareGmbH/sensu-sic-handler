// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	"github.com/mattn/go-redmine"
)

type projectMemberships struct {
	project     redmine.Project
	memberships map[int][]redmine.IdName
}
