package media

import "context"

type CommandRunner interface {
	Run(context.Context, string, ...string) ([]byte, error)
}

type MPRIS struct{ runner CommandRunner }

func NewMPRIS(runner CommandRunner) MPRIS { return MPRIS{runner: runner} }
func (m MPRIS) PlayPause(ctx context.Context) error {
	_, err := m.runner.Run(ctx, "playerctl", "play-pause")
	return err
}
func (m MPRIS) Next(ctx context.Context) error {
	_, err := m.runner.Run(ctx, "playerctl", "next")
	return err
}
func (m MPRIS) Previous(ctx context.Context) error {
	_, err := m.runner.Run(ctx, "playerctl", "previous")
	return err
}
func (m MPRIS) SetVolume(ctx context.Context, value string) error {
	_, err := m.runner.Run(ctx, "playerctl", "volume", value)
	return err
}
