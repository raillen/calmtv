package games

import (
	"strings"
	"testing"
)

func TestClassifyUsesExtensionAndStableHash(t *testing.T) {
	rom, err := Classify("Super Mario.nes", strings.NewReader("fixture"))
	if err != nil {
		t.Fatal(err)
	}
	if rom.System != NES || rom.Core != "mesen" || len(rom.Hash) != 64 {
		t.Fatalf("rom = %#v", rom)
	}
}
