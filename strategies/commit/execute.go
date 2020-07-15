package commit

import (
	"fmt"

	"github.com/google/go-github/github"
)

func (dc DirectCommit) Execute() error {
	content, err := dc.comment.FileContent()
	if err != nil {
		return err
	}
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String(fmt.Sprintf("Add %s comment to post %s", dc.comment.AuthorName, dc.comment.Resource)),
		Content:   content,
		Branch:    &dc.repository.Branch,
		Committer: &github.CommitAuthor{Name: &dc.comment.AuthorName, Email: &dc.comment.AuthorEmail},
	}
	_, _, err = dc.repository.GitHubClient.Repositories.CreateFile(dc.ctx, dc.repository.Owner, dc.repository.Name, dc.comment.FileName(), opts)
	if err != nil {
		return fmt.Errorf("an error occurred while posting the comment: %s", err)
	}
	return nil
}
