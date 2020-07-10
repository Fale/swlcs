package commit

import (
	"context"

	"github.com/fale/swlcs/objects"
)

type DirectCommit struct {
	repository *objects.Repository
	comment    *objects.Comment
	ctx        context.Context
}

func Init(ctx context.Context, repository *objects.Repository, comment *objects.Comment) DirectCommit {
	return DirectCommit{
		repository: repository,
		comment:    comment,
		ctx:        ctx,
	}
}
