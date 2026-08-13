package system

import (
	"context"
	"strconv"
)

// Facade exposes product-level system operations to screens. It keeps CLI and
// D-Bus details inside adapters and gives the UI one place for errors/status.
type Facade struct {
	Network   NetworkAdapter
	Bluetooth BluetoothAdapter
	Audio     AudioAdapter
	Display   DisplayAdapter
	Storage   StorageAdapter
	Power     PowerAdapter
}

func NewFacade(runner Runner) Facade {
	return Facade{
		Network: NewNetworkAdapter(runner), Bluetooth: NewBluetoothAdapter(runner), Audio: NewAudioAdapter(runner), Display: NewDisplayAdapter(runner), Storage: NewStorageAdapter(runner), Power: NewPowerAdapter(runner),
	}
}

type QuickSettingResult struct{ Summary string }

func (f Facade) Execute(ctx context.Context, action string) (QuickSettingResult, error) {
	switch action {
	case "wifi":
		items, err := f.Network.List(ctx)
		return QuickSettingResult{Summary: "Wi-Fi: " + count(items)}, err
	case "bluetooth":
		items, err := f.Bluetooth.List(ctx)
		return QuickSettingResult{Summary: "Bluetooth: " + count(items)}, err
	case "volume":
		return QuickSettingResult{Summary: "Volume disponível"}, f.Audio.SetVolume(ctx, "@DEFAULT_AUDIO_SINK@", 50)
	case "audio-output":
		items, err := f.Audio.ListOutputs(ctx)
		return QuickSettingResult{Summary: "Saídas: " + count(items)}, err
	case "display":
		items, err := f.Display.ListModes(ctx)
		return QuickSettingResult{Summary: "Modos: " + count(items)}, err
	case "storage", "usb":
		items, err := f.Storage.ListRemovable(ctx)
		return QuickSettingResult{Summary: "Volumes: " + count(items)}, err
	case "suspend":
		return QuickSettingResult{Summary: "Suspensão solicitada"}, f.Power.Suspend(ctx)
	case "reboot":
		return QuickSettingResult{Summary: "Reinicialização solicitada"}, f.Power.Reboot(ctx)
	case "power-off":
		return QuickSettingResult{Summary: "Desligamento solicitado"}, f.Power.PowerOff(ctx)
	default:
		return QuickSettingResult{Summary: action}, nil
	}
}

func count[T any](items []T) string {
	if len(items) == 1 {
		return "1 item"
	}
	return stringCount(len(items)) + " itens"
}

func stringCount(value int) string {
	if value == 0 {
		return "nenhum"
	}
	return strconv.Itoa(value)
}
