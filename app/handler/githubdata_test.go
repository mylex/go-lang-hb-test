package handler

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func TestInitClient(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		client = github.NewClient(nil)
	} else {
		tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		))
		client = github.NewClient(tc)
	}

	output := InitClient()

	if !reflect.DeepEqual(output, client) {
		t.Error("Expected Github Client")
	}
}

func TestFetchFollowersList(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"mylex"},
		{"yasirtaher"},
		{"marufhub"},
		{"undefined"},
		{"null"},
	}

	for _, test := range tests {
		results, err := fetchFollowersList(test.input)
		if err != nil {
			for _, result := range results {
				if result.GetLogin() == "" {
					t.Error("Expected github username")
				}
			}
		}
	}
}

func TestFetchStarCountOfRepository(t *testing.T) {
	var tests = []struct {
		name      string
		repo_name string
	}{
		{"mylex", "LuckyWinner-Flutter"},
		{"yasirtaher", "undefined"},
		{"marufhub", "null"},
		{"undefined", "undefined"},
		{"null", "null"},
	}

	for _, test := range tests {
		repositry, _ := fetchStarCountOfRepository(test.name, test.repo_name)

		if repositry.GetStargazersCount() < 0 {
			t.Error("Expected al least 0")
		}
	}
}
