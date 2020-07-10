package objects

import "github.com/google/go-github/github"

type Repository struct {
	GitHubClient *github.Client
	Owner        string
	Name         string
	Branch       string
}
