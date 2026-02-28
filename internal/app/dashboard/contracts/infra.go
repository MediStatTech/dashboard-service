package contracts

import (
	"context"

	"github.com/MediStatTech/commitplan"
)

// ============================================================================
// Infrastructure Interfaces
// ============================================================================

type Logger interface {
	Debug(msg string, fields map[string]any)
	Info(msg string, fields map[string]any)
	Warn(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
	Fatal(msg string, fields map[string]any)
}

type Committer interface {
	Apply(ctx context.Context, p commitplan.PlanLike, opts ...commitplan.ApplyOption) error
}

type Plan interface {
	Mutations() []Mutation
	AddMut(mutation Mutation)
	AddMuts(mutations ...Mutation)
}

type Mutation interface {
	Query() string
	Args() []any
}
