package pr

import (
	"context"
	"fmt"

	"github.com/fale/swlcs/objects"
)

type PR struct {
	repository   *objects.Repository
	comment      *objects.Comment
	ctx          context.Context
	subject      string
	commitBranch string
}

func Init(ctx context.Context, repository *objects.Repository, comment *objects.Comment) PR {
	return PR{
		repository:   repository,
		comment:      comment,
		ctx:          ctx,
		subject:      fmt.Sprintf("Add comment to %s", comment.Resource),
		commitBranch: fmt.Sprintf("%s-%s", comment.Resource, comment.Time.Format("2006-01-02_15-04-05")),
	}
}
