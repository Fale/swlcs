package commit

import (
	"fmt"

	"github.com/google/go-github/github"
)

func (dc DirectCommit) Execute() error {
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String(fmt.Sprintf("Add %s comment to post %s", dc.comment.AuthorName, dc.comment.Resource)),
		Content:   []byte(dc.comment.Content),
		Branch:    &dc.repository.Branch,
		Committer: &github.CommitAuthor{Name: &dc.comment.AuthorName, Email: &dc.comment.AuthorEmail},
	}
	fileName := fmt.Sprintf("data/comments/post/%s/%s.md", dc.comment.Resource, dc.comment.Time.Format("2006-01-02_15-04-05"))
	_, _, err := dc.repository.GitHubClient.Repositories.CreateFile(dc.ctx, dc.repository.Owner, dc.repository.Name, fileName, opts)
	if err != nil {
		return fmt.Errorf("an error occurred while posting the comment: %s", err)
	}
	return nil
}
