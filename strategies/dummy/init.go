package dummy

import (
	"context"

	"github.com/fale/swlcs/objects"
)

type Dummy struct {
	repository *objects.Repository
	comment    *objects.Comment
	ctx        context.Context
}

func Init(ctx context.Context, repository *objects.Repository, comment *objects.Comment) Dummy {
	return Dummy{
		repository: repository,
		comment:    comment,
		ctx:        ctx,
	}
}
