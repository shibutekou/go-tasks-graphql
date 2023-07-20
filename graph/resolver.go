package graph

import (
	"github.com/vgekko/go-tasks-graphql/internal/usecase"
	"golang.org/x/exp/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Uc  *usecase.Usecase
	Log *slog.Logger
}
