package system

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type NetworkAdapter struct{ runner Runner }
type BluetoothAdapter struct{ runner Runner }
type AudioAdapter struct{ runner Runner }
type DisplayAdapter struct{ runner Runner }
type StorageAdapter struct{ runner Runner }
type PowerAdapter struct{ runner Runner }

func NewNetworkAdapter(runner Runner) NetworkAdapter     { return NetworkAdapter{runner: runner} }
func NewBluetoothAdapter(runner Runner) BluetoothAdapter { return BluetoothAdapter{runner: runner} }
func NewAudioAdapter(runner Runner) AudioAdapter         { return AudioAdapter{runner: runner} }
func NewDisplayAdapter(runner Runner) DisplayAdapter     { return DisplayAdapter{runner: runner} }
func NewStorageAdapter(runner Runner) StorageAdapter     { return StorageAdapter{runner: runner} }
func NewPowerAdapter(runner Runner) PowerAdapter         { return PowerAdapter{runner: runner} }

type Network struct {
	Name   string
	Device string
	Active bool
}

func (a NetworkAdapter) List(ctx context.Context) ([]Network, error) {
	output, err := a.runner.Run(ctx, "nmcli", "-t", "-f", "NAME,DEVICE,ACTIVE", "connection", "show")
	if err != nil {
		return nil, translateError("listar redes", err)
	}
	var networks []Network
	for _, line := range nonEmptyLines(string(output)) {
		fields := strings.Split(line, ":")
		if len(fields) < 3 {
			continue
		}
		networks = append(networks, Network{Name: fields[0], Device: fields[1], Active: fields[2] == "yes"})
	}
	return networks, nil
}

func (a NetworkAdapter) Connect(ctx context.Context, name string) error {
	if strings.TrimSpace(name) == "" {
		return &ServiceError{Code: ErrorInvalid, Message: "A rede selecionada é inválida."}
	}
	_, err := a.runner.Run(ctx, "nmcli", "connection", "up", "id", name)
	return translateError("conectar à rede", err)
}

type BluetoothDevice struct {
	Address   string
	Name      string
	Paired    bool
	Connected bool
}

func (a BluetoothAdapter) List(ctx context.Context) ([]BluetoothDevice, error) {
	output, err := a.runner.Run(ctx, "bluetoothctl", "devices")
	if err != nil {
		return nil, translateError("listar dispositivos Bluetooth", err)
	}
	var devices []BluetoothDevice
	for _, line := range nonEmptyLines(string(output)) {
		fields := strings.SplitN(line, " ", 3)
		if len(fields) == 3 && fields[0] == "Device" {
			devices = append(devices, BluetoothDevice{Address: fields[1], Name: fields[2]})
		}
	}
	return devices, nil
}

func (a BluetoothAdapter) SetPowered(ctx context.Context, enabled bool) error {
	value := "off"
	if enabled {
		value = "on"
	}
	_, err := a.runner.Run(ctx, "bluetoothctl", "power", value)
	return translateError("alterar Bluetooth", err)
}

type AudioOutput struct{ ID, Name string }

func (a AudioAdapter) ListOutputs(ctx context.Context) ([]AudioOutput, error) {
	output, err := a.runner.Run(ctx, "wpctl", "status")
	if err != nil {
		return nil, translateError("listar saídas de áudio", err)
	}
	var outputs []AudioOutput
	for _, line := range nonEmptyLines(string(output)) {
		trimmed := strings.TrimSpace(line)
		if !strings.Contains(trimmed, ".") || !strings.Contains(trimmed, "[vol:") {
			continue
		}
		parts := strings.SplitN(trimmed, ".", 2)
		if len(parts) != 2 {
			continue
		}
		id := strings.TrimSpace(parts[0])
		if _, parseErr := strconv.Atoi(id); parseErr != nil {
			continue
		}
		outputs = append(outputs, AudioOutput{ID: id, Name: strings.TrimSpace(strings.Split(parts[1], "[")[0])})
	}
	return outputs, nil
}

func (a AudioAdapter) SetVolume(ctx context.Context, outputID string, percent int) error {
	if percent < 0 || percent > 100 || strings.TrimSpace(outputID) == "" {
		return &ServiceError{Code: ErrorInvalid, Message: "O volume selecionado é inválido."}
	}
	_, err := a.runner.Run(ctx, "wpctl", "set-volume", outputID, fmt.Sprintf("%d%%", percent))
	return translateError("alterar volume", err)
}

func (a AudioAdapter) SetMute(ctx context.Context, outputID string, muted bool) error {
	if strings.TrimSpace(outputID) == "" {
		return &ServiceError{Code: ErrorInvalid, Message: "A saída de áudio selecionada é inválida."}
	}
	value := "0"
	if muted {
		value = "1"
	}
	_, err := a.runner.Run(ctx, "wpctl", "set-mute", outputID, value)
	return translateError("alterar silêncio do áudio", err)
}

type DisplayMode struct{ Output, Mode string }

func (a DisplayAdapter) ListModes(ctx context.Context) ([]DisplayMode, error) {
	output, err := a.runner.Run(ctx, "xrandr", "--query")
	if err != nil {
		return nil, translateError("listar displays", err)
	}
	var modes []DisplayMode
	currentOutput := ""
	for _, line := range nonEmptyLines(string(output)) {
		if !strings.HasPrefix(line, " ") && strings.Contains(line, " connected") {
			currentOutput = strings.Fields(line)[0]
			continue
		}
		trimmed := strings.TrimSpace(line)
		if currentOutput != "" && len(trimmed) > 0 && trimmed[0] >= '0' && trimmed[0] <= '9' {
			modes = append(modes, DisplayMode{Output: currentOutput, Mode: strings.Fields(trimmed)[0]})
		}
	}
	return modes, nil
}

func (a DisplayAdapter) SetMode(ctx context.Context, output, mode string) error {
	if strings.TrimSpace(output) == "" || strings.TrimSpace(mode) == "" || strings.ContainsAny(output+mode, ";&|`$\n") {
		return &ServiceError{Code: ErrorInvalid, Message: "O modo de display selecionado é inválido."}
	}
	_, err := a.runner.Run(ctx, "xrandr", "--output", output, "--mode", mode)
	return translateError("alterar resolução", err)
}

type Volume struct{ Device, Label, MountPoint string }

func (a StorageAdapter) ListRemovable(ctx context.Context) ([]Volume, error) {
	output, err := a.runner.Run(ctx, "udisksctl", "status")
	if err != nil {
		return nil, translateError("listar armazenamento USB", err)
	}
	var volumes []Volume
	for _, line := range nonEmptyLines(string(output)) {
		fields := strings.Fields(line)
		if len(fields) >= 2 && strings.HasPrefix(fields[0], "/dev/") {
			volumes = append(volumes, Volume{Device: fields[0], Label: strings.Join(fields[1:], " ")})
		}
	}
	return volumes, nil
}

func (a StorageAdapter) Mount(ctx context.Context, device string) error {
	if strings.TrimSpace(device) == "" || !strings.HasPrefix(device, "/dev/") || strings.ContainsAny(device, ";&|`$\n") {
		return &ServiceError{Code: ErrorInvalid, Message: "O dispositivo selecionado é inválido."}
	}
	_, err := a.runner.Run(ctx, "udisksctl", "mount", "-b", device)
	return translateError("montar armazenamento", err)
}

func (a PowerAdapter) Suspend(ctx context.Context) error {
	return a.run(ctx, "suspender o sistema", "suspend")
}
func (a PowerAdapter) Reboot(ctx context.Context) error {
	return a.run(ctx, "reiniciar o sistema", "reboot")
}
func (a PowerAdapter) PowerOff(ctx context.Context) error {
	return a.run(ctx, "desligar o sistema", "poweroff")
}

func (a PowerAdapter) run(ctx context.Context, action, command string) error {
	_, err := a.runner.Run(ctx, "loginctl", command)
	return translateError(action, err)
}

func nonEmptyLines(value string) []string {
	var lines []string
	for _, line := range strings.Split(value, "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
