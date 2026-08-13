package input

import "testing"

func TestFromKeyMapsRemoteActions(t *testing.T) {
	tests := map[string]Action{
		"Up":            NavUp,
		"Return":        Accept,
		"Escape":        Back,
		"Home":          Home,
		"XF86AudioMute": Mute,
		"Page_Down":     ChannelDown,
	}

	for key, want := range tests {
		got, ok := FromKey(key)
		if !ok || got != want {
			t.Fatalf("FromKey(%q) = %q, %v; want %q, true", key, got, ok, want)
		}
	}
}

func TestFromKeyRejectsUnknownKey(t *testing.T) {
	if _, ok := FromKey("F13"); ok {
		t.Fatal("unknown key should not produce an action")
	}
}

func TestFromKeyMapsChannelNumbers(t *testing.T) {
	action, ok := FromKey("7")
	if !ok || action != Action("CHANNEL_7") {
		t.Fatalf("action=%q ok=%v", action, ok)
	}
	number, ok := ChannelNumber(action)
	if !ok || number != 7 {
		t.Fatalf("number=%d ok=%v", number, ok)
	}
}
