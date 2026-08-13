package updates

import (
	"context"
	"fmt"
)

type Runner interface {
	Run(context.Context, string, ...string) error
}

type Service struct{ runner Runner }

func NewService(runner Runner) Service { return Service{runner: runner} }

// Update delegates privilege and policy to the installed system helper. The
// Shell never calls sudo and never accepts arbitrary package names from UI.
func (s Service) Update(ctx context.Context) error {
	if err := s.runner.Run(ctx, "tv-shell-update"); err != nil {
		return fmt.Errorf("atualizar Calm TV: %w", err)
	}
	return nil
}
