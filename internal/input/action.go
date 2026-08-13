package input

import (
	"strconv"
	"strings"
)

// Action is a semantic command emitted by every input backend.
type Action string

const (
	NavUp       Action = "NAV_UP"
	NavDown     Action = "NAV_DOWN"
	NavLeft     Action = "NAV_LEFT"
	NavRight    Action = "NAV_RIGHT"
	Accept      Action = "ACCEPT"
	Back        Action = "BACK"
	Home        Action = "HOME"
	Menu        Action = "MENU"
	Search      Action = "SEARCH"
	PlayPause   Action = "PLAY_PAUSE"
	Next        Action = "NEXT"
	Previous    Action = "PREVIOUS"
	VolUp       Action = "VOL_UP"
	VolDown     Action = "VOL_DOWN"
	Mute        Action = "MUTE"
	ChannelUp   Action = "CHANNEL_UP"
	ChannelDown Action = "CHANNEL_DOWN"
)

// Event keeps physical-source information available for diagnostics without
// leaking device-specific details into screen code.
type Event struct {
	Action Action
	Source string
}

var keyActions = map[string]Action{
	"Up":                   NavUp,
	"Down":                 NavDown,
	"Left":                 NavLeft,
	"Right":                NavRight,
	"Return":               Accept,
	"KP_Enter":             Accept,
	"Escape":               Back,
	"BackSpace":            Back,
	"Home":                 Home,
	"h":                    Home,
	"H":                    Home,
	"Menu":                 Menu,
	"m":                    Menu,
	"M":                    Menu,
	"slash":                Search,
	"question":             Search,
	"space":                PlayPause,
	"XF86AudioPlay":        PlayPause,
	"XF86AudioPause":       PlayPause,
	"XF86AudioRaiseVolume": VolUp,
	"XF86AudioLowerVolume": VolDown,
	"XF86AudioMute":        Mute,
	"XF86AudioNext":        Next,
	"XF86AudioPrev":        Previous,
	"Page_Up":              ChannelUp,
	"Page_Down":            ChannelDown,
}

func init() {
	for digit := 0; digit <= 9; digit++ {
		action := Action("CHANNEL_" + strconv.Itoa(digit))
		key := strconv.Itoa(digit)
		keyActions[key] = action
		keyActions["KP_"+key] = action
	}
}

// FromKey maps a GTK/GDK key name to a semantic action.
func FromKey(keyName string) (Action, bool) {
	action, ok := keyActions[keyName]
	return action, ok
}

func ChannelNumber(action Action) (int, bool) {
	value := strings.TrimPrefix(string(action), "CHANNEL_")
	if value == string(action) {
		return 0, false
	}
	number, err := strconv.Atoi(value)
	if err != nil || number < 0 || number > 9 {
		return 0, false
	}
	return number, true
}
