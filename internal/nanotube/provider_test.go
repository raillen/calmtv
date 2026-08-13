package nanotube

import (
	"context"
	"testing"
)

func TestUnavailableProviderFailsWithoutStartingBackgroundProcess(t *testing.T) {
	_, err := (UnavailableProvider{}).Search(context.Background(), "movie", 10)
	if err != ErrUnavailable {
		t.Fatalf("error = %v", err)
	}
}
