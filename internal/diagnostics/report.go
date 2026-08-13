package diagnostics

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Command interface {
	Run(ctx context.Context, name string, args ...string) ([]byte, error)
}

type Reporter struct{ command Command }

type Report struct {
	OS       string            `json:"os"`
	Arch     string            `json:"arch"`
	CPU      int               `json:"cpu"`
	RAM      string            `json:"ram"`
	Commands map[string]string `json:"commands"`
}

func NewReporter(command Command) Reporter { return Reporter{command: command} }

func (r Reporter) Collect(ctx context.Context) Report {
	report := Report{OS: runtime.GOOS, Arch: runtime.GOARCH, CPU: runtime.NumCPU(), RAM: os.Getenv("TV_SHELL_RAM"), Commands: map[string]string{}}
	for name, args := range map[string][]string{
		"kernel":    {"uname", "-sr"},
		"ram":       {"cat", "/proc/meminfo"},
		"disk":      {"df", "-h", "/"},
		"zram":      {"zramctl"},
		"network":   {"nmcli", "general", "status"},
		"wifi":      {"nmcli", "radio", "wifi"},
		"bluetooth": {"bluetoothctl", "show"},
		"pipewire":  {"wpctl", "status"},
		"display":   {"xrandr", "--query"},
		"va_api":    {"vainfo"},
		"mpv":       {"mpv", "--version"},
		"input":     {"libinput", "list-devices"},
		"systemd":   {"systemctl", "--user", "--failed"},
		"cgroups":   {"systemctl", "--user", "status"},
	} {
		if len(args) == 0 {
			continue
		}
		output, err := r.command.Run(ctx, args[0], args[1:]...)
		if err != nil {
			report.Commands[name] = "unavailable"
			continue
		}
		report.Commands[name] = redact(string(output))
	}
	return report
}

func Save(path string, report Report) error {
	if path == "" || filepath.Base(path) != "latest.json" {
		return fmt.Errorf("diagnostic report path must be latest.json")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return err
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("encode diagnostic report: %w", err)
	}
	temporary := path + ".tmp"
	if err := os.WriteFile(temporary, append(data, '\n'), 0600); err != nil {
		return err
	}
	if err := os.Rename(temporary, path); err != nil {
		_ = os.Remove(temporary)
		return err
	}
	return nil
}

type ExecCommand struct{}

func (ExecCommand) Run(ctx context.Context, name string, args ...string) ([]byte, error) {
	return exec.CommandContext(ctx, name, args...).CombinedOutput()
}

func redact(value string) string {
	for _, key := range []string{"token", "password", "secret", "cookie"} {
		value = redactKey(value, key)
	}
	return strings.TrimSpace(value)
}

func redactKey(value, key string) string {
	lines := strings.Split(value, "\n")
	for index, line := range lines {
		if strings.Contains(strings.ToLower(line), key) {
			lines[index] = fmt.Sprintf("[%s redacted]", key)
		}
	}
	return strings.Join(lines, "\n")
}
