package objects

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

// Comment is the structure that contains a comment.
type Comment struct {
	// Resource is the unique identifier of the blog post or page the comment refers to
	Resource string
	// AuthorName is the name of the author of the comment
	AuthorName string
	// AuthorEmail is the email of the author of the comment
	AuthorEmail string
	// Body is the text of the comment
	Body string
	// Time is the date/time the comment was created
	Time time.Time
}

// FileContent returns the comment in YAML format
func (c Comment) FileContent() ([]byte, error) {
	d := map[string]string{
		"date":  c.Time.Format("2006-01-02 15:04:05"),
		"name":  c.AuthorName,
		"email": c.AuthorEmail,
		"body":  c.Body,
	}
	comment, err := yaml.Marshal(&d)
	if err != nil {
		return nil, err
	}
	var out []byte
	out = append(out, "---\n"...)
	out = append(out, comment...)
	out = append(out, "---"...)
	return out, nil
}

// FileName returns the name that the file will need to have to properly read by the rendering function
func (c Comment) FileName() string {
	return fmt.Sprintf("data/comments/%s/%s.yaml", c.Resource, c.Time.Format("2006-01-02_15-04-05"))
}
