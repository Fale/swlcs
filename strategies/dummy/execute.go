package dummy

import (
	"fmt"
)

func (d Dummy) Execute() error {
	content, err := d.comment.FileContent()
	if err != nil {
		return err
	}
	fmt.Printf("Commit Message: Add %s comment to post %s\n", d.comment.AuthorName, d.comment.Resource)
	fmt.Printf("Repository owner/name: %s/%s\n", d.repository.Owner, d.repository.Name)
	fmt.Printf("Repository branch: %s\n", d.repository.Branch)
	fmt.Printf("Author: %s <%s>\n", d.comment.AuthorName, d.comment.AuthorEmail)
	fmt.Printf("File name: %s\n", d.comment.FileName())
	fmt.Printf("File content:\n%s\n", content)
	return nil
}
