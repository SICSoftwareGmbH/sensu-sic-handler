// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	redmine "github.com/mattn/go-redmine"
)

type projectMemberships struct {
	project     redmine.Project
	memberships map[int][]redmine.IdName
}
