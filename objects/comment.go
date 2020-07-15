package objects

import "time"

type Comment struct {
	Resource    string
	AuthorName  string
	AuthorEmail string
	FileName    string
	Content     string
	Time        time.Time
}
