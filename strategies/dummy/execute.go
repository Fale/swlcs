package dummy

import (
	"log"
)

func (d Dummy) Execute() error {
	content, err := d.comment.FileContent()
	if err != nil {
		return err
	}
	log.Printf("Commit Message: Add %s comment to post %s\n", d.comment.AuthorName, d.comment.Resource)
	log.Printf("Repository owner/name: %s/%s\n", d.repository.Owner, d.repository.Name)
	log.Printf("Repository branch: %s\n", d.repository.Branch)
	log.Printf("Author: %s <%s>\n", d.comment.AuthorName, d.comment.AuthorEmail)
	log.Printf("File name: %s\n", d.comment.FileName())
	log.Printf("File content:\n%s\n", content)
	return nil
}
