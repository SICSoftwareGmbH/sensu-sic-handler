// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mattn/go-redmine"
	etcd "go.etcd.io/etcd/clientv3"
)

// Import imports memberships and users into etcd
func Import(url string, token string, etcdClient *etcd.Client) error {
	projectMemberships, users, err := loadRedmineData(url, token)
	if err != nil {
		return err
	}

	err = writeToEtcd(etcdClient, projectMemberships, users)
	if err != nil {
		return err
	}

	return nil
}

func loadRedmineData(url string, token string) ([]projectMemberships, map[int]redmine.User, error) {
	client := redmine.NewClient(url, token)

	projects, err := redmineProjects(client)
	if err != nil {
		return nil, nil, err
	}

	memberships, err := redmineProjectsMemberships(client, projects)
	if err != nil {
		return nil, nil, err
	}

	users, err := redmineUsers(client, memberships)
	if err != nil {
		return nil, nil, err
	}

	return memberships, users, nil
}

func redmineProjects(client *redmine.Client) ([]redmine.Project, error) {
	projects := make([]redmine.Project, 0)
	lastCount := 1

	client.Offset = 0
	client.Limit = 100

	for lastCount > 0 {
		items, err := client.Projects()
		if err != nil {
			return nil, fmt.Errorf("unable to fetch projects: %q", err)
		}

		projects = append(projects, items...)

		client.Offset += client.Limit

		lastCount = len(items)
	}

	return projects, nil
}

func redmineProjectsMemberships(client *redmine.Client, projects []redmine.Project) ([]projectMemberships, error) {
	client.Offset = 0
	client.Limit = 100

	memberships := make([]projectMemberships, 0)

	for _, project := range projects {
		pm := make(map[int][]redmine.IdName)

		items, err := client.Memberships(project.Id)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch memberships for project %s (%d): %q", project.Identifier, project.Id, err)
		}

		for _, m := range items {
			for _, r := range m.Roles {
				// users list may contain invalid users
				if m.User.Id == 0 {
					continue
				}

				pm[r.Id] = append(pm[r.Id], m.User)
			}
		}

		memberships = append(memberships, projectMemberships{project: project, memberships: pm})
	}

	return memberships, nil
}

func redmineUsers(client *redmine.Client, projectMemberships []projectMemberships) (map[int]redmine.User, error) {
	client.Offset = 0
	client.Limit = 100

	usersMap := make(map[int]redmine.User)

	for _, projectMembership := range projectMemberships {
		for _, users := range projectMembership.memberships {
			for _, user := range users {
				if _, ok := usersMap[user.Id]; !ok {
					u, err := client.User(user.Id)
					if err != nil {
						return nil, fmt.Errorf("unable to fetch user %s (%d): %q", user.Name, user.Id, err)
					}

					usersMap[user.Id] = *u
				}
			}
		}
	}

	return usersMap, nil
}

func writeToEtcd(client *etcd.Client, projectMemberships []projectMemberships, usersMap map[int]redmine.User) error {
	for _, projectMembership := range projectMemberships {
		for roleID, users := range projectMembership.memberships {
			mails := make([]string, 0)

			for _, user := range users {
				mails = append(mails, usersMap[user.Id].Mail)
			}

			data, err := json.Marshal(mails)
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_, err = client.Put(ctx, EtcdKey(projectMembership.project.Identifier, roleID, "mail"), string(data))
			cancel()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
