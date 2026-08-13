package games

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"path/filepath"
	"strings"
)

type System string

const (
	NES          System = "nes"
	SNES         System = "snes"
	GB           System = "game-boy"
	GBC          System = "game-boy-color"
	GBA          System = "game-boy-advance"
	MegaDrive    System = "mega-drive"
	MasterSystem System = "master-system"
	GameGear     System = "game-gear"
)

type ROM struct {
	Path, Hash, Title string
	System            System
	Core              string
}

var extensions = map[string]System{
	".nes": NES, ".smc": SNES, ".sfc": SNES,
	".gb": GB, ".gbc": GBC, ".gba": GBA,
	".md": MegaDrive, ".gen": MegaDrive, ".sms": MasterSystem, ".gg": GameGear,
}

func Classify(path string, reader io.Reader) (ROM, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return ROM{}, err
	}
	romSystem := extensions[strings.ToLower(filepath.Ext(path))]
	return ROM{Path: path, Hash: hex.EncodeToString(hash.Sum(nil)), Title: strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)), System: romSystem, Core: coreFor(romSystem)}, nil
}

func coreFor(system System) string {
	switch system {
	case NES:
		return "mesen"
	case SNES:
		return "snes9x"
	case GB, GBC:
		return "gambatte"
	case GBA:
		return "mgba"
	case MegaDrive, MasterSystem, GameGear:
		return "genesis_plus_gx"
	default:
		return ""
	}
}
