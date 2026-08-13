package shell

import (
	"os"
	"strconv"

	"github.com/gotk3/gotk3/gdk"
)

// resolveMonitorIndex keeps the selected output inside the monitor list and
// falls back to the primary output when configuration is absent or invalid.
func resolveMonitorIndex(value string, monitorCount, primary int) int {
	if monitorCount <= 0 {
		return 0
	}
	if primary < 0 || primary >= monitorCount {
		primary = 0
	}
	if value == "" || value == "primary" {
		return primary
	}
	index, err := strconv.Atoi(value)
	if err != nil || index < 0 || index >= monitorCount {
		return primary
	}
	return index
}

func (a *App) fullscreenOnConfiguredMonitor() {
	display, err := gdk.DisplayGetDefault()
	if err != nil {
		a.window.Fullscreen()
		return
	}
	screen, err := display.GetDefaultScreen()
	if err != nil {
		a.window.Fullscreen()
		return
	}

	monitorCount := display.GetNMonitors()
	if monitorCount <= 0 {
		a.window.Fullscreen()
		return
	}
	primary := 0
	if primaryMonitor, primaryErr := display.GetPrimaryMonitor(); primaryErr == nil && primaryMonitor != nil {
		for index := 0; index < monitorCount; index++ {
			monitor, monitorErr := display.GetMonitor(index)
			if monitorErr == nil && monitor.Native() == primaryMonitor.Native() {
				primary = index
				break
			}
		}
	}
	monitor := resolveMonitorIndex(os.Getenv("TV_SHELL_MONITOR"), monitorCount, primary)
	a.window.FullscreenOnMonitor(screen, monitor)
}
