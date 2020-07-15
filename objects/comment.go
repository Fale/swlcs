package objects

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type Comment struct {
	Resource    string
	AuthorName  string
	AuthorEmail string
	Body        string
	Time        time.Time
}

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

func (c Comment) FileName() string {
	return fmt.Sprintf("data/comments/%s/%s.yaml", c.Resource, c.Time.Format("2006-01-02_15-04-05"))
}
