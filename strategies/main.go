package strategies

import (
	"context"
	"fmt"

	"github.com/fale/swlcs/objects"
	"github.com/fale/swlcs/strategies/commit"
	"github.com/fale/swlcs/strategies/dummy"
	"github.com/fale/swlcs/strategies/pr"
)

type Strategy interface {
	Execute() error
}

func Init(strategy string, ctx context.Context, repository *objects.Repository, comment *objects.Comment) (Strategy, error) {
	switch strategy {
	case "commit":
		return commit.Init(ctx, repository, comment), nil
	case "dummy":
		return dummy.Init(ctx, repository, comment), nil
	case "pr":
		return pr.Init(ctx, repository, comment), nil
	default:
		return nil, fmt.Errorf("strategy not found")
	}
}
