package handler

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func InitClient() *github.Client {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return github.NewClient(nil)
	} else {
		tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		))
		return github.NewClient(tc)
	}
}

// Fetch all the followers List of a github user.
func fetchFollowersList(username string) ([]*github.User, error) {
	client := InitClient()
	orgs, _, err := client.Users.ListFollowers(context.Background(), username, nil)
	return orgs, err
}

// Fetch stars count of a github repository.
func fetchStarCountOfRepository(username string, repositoryName string) (*github.Repository, error) {
	client := InitClient()
	orgs, _, err := client.Repositories.Get(context.Background(), username, repositoryName)
	return orgs, err
}
