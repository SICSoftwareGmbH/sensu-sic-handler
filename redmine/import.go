// Copyright © 2019 SIC! Software GmbH

package redmine

import (
	"encoding/json"
	"errors"

	"github.com/go-redis/redis"
	redmine "github.com/mattn/go-redmine"
)

// Import imports memberships and users into redis
func Import(url string, token string, redisClient *redis.Client) error {
	projectMemberships, users, err := loadRedmineData(url, token)
	if err != nil {
		return err
	}

	err = writeToRedis(redisClient, projectMemberships, users)
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
			return nil, errors.New("unable to fetch projects")
		}

		for _, p := range items {
			projects = append(projects, p)
		}

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
		pm := make(map[int][]redmine.IdName, 0)

		items, err := client.Memberships(project.Id)
		if err != nil {
			return nil, errors.New("unable to fetch project memberships")
		}

		for _, m := range items {
			for _, r := range m.Roles {
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
						return nil, errors.New("unable to fetch user")
					}

					usersMap[user.Id] = *u
				}
			}
		}
	}

	return usersMap, nil
}

func writeToRedis(client *redis.Client, projectMemberships []projectMemberships, usersMap map[int]redmine.User) error {
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

			err = client.Set(RedisKey(projectMembership.project.Identifier, roleID, "mail"), data, 0).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
