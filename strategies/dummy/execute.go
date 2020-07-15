package dummy

import (
	"fmt"
)

func (d Dummy) Execute() error {
	fileName := fmt.Sprintf("data/comments/post/%s/%s.md", d.comment.Resource, d.comment.Time.Format("2006-01-02_15-04-05"))
	fmt.Printf("Commit Message: Add %s comment to post %s\n", d.comment.AuthorName, d.comment.Resource)
	fmt.Printf("Repository owner/name: %s/%s\n", d.repository.Owner, d.repository.Name)
	fmt.Printf("Repository branch: %s\n", d.repository.Branch)
	fmt.Printf("Author: %s <%s>\n", d.comment.AuthorName, d.comment.AuthorEmail)
	fmt.Printf("File name: %s\n", fileName)
	fmt.Printf("File content:\n%s\n", d.comment.Content)
	return nil
}
