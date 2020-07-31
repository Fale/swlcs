package objects

import "github.com/google/go-github/github"

// Repository is the structure that contains the information of the GitHub repository where the website is hosted
type Repository struct {
	// GitHubClient is the reference to a GitHub client instance
	GitHubClient *github.Client
	// Owner is the name of the repository Owner on GitHub
	Owner string
	// Name is the name of the repository on GitHub
	Name string
	// Branch is the name of the branch that contains the website
	Branch string
}
