package nanotube

import (
	"context"
	"time"
)

type Content struct {
	ID, Title, URL, Thumbnail string
	Duration                  time.Duration
}

// Provider is the only Shell-facing NanoTube boundary. Resolver/auth/catalog
// implementation stays in the separate NanoTube process or package.
type Provider interface {
	ContinueWatching(ctx context.Context, limit int) ([]Content, error)
	Search(ctx context.Context, query string, limit int) ([]Content, error)
	Open(ctx context.Context, contentID string) (Content, error)
}

type UnavailableProvider struct{}

func (UnavailableProvider) ContinueWatching(context.Context, int) ([]Content, error) {
	return nil, ErrUnavailable
}
func (UnavailableProvider) Search(context.Context, string, int) ([]Content, error) {
	return nil, ErrUnavailable
}
func (UnavailableProvider) Open(context.Context, string) (Content, error) {
	return Content{}, ErrUnavailable
}

var ErrUnavailable = unavailableError{}

type unavailableError struct{}

func (unavailableError) Error() string { return "NanoTube indisponível" }
