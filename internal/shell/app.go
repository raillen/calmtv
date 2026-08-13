package shell

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/raillen/calmtv/internal/appmanager"
	"github.com/raillen/calmtv/internal/diagnostics"
	"github.com/raillen/calmtv/internal/focus"
	"github.com/raillen/calmtv/internal/games"
	"github.com/raillen/calmtv/internal/input"
	"github.com/raillen/calmtv/internal/iptv"
	"github.com/raillen/calmtv/internal/library"
	"github.com/raillen/calmtv/internal/media"
	"github.com/raillen/calmtv/internal/recovery"
	"github.com/raillen/calmtv/internal/state"
	"github.com/raillen/calmtv/internal/system"
	"github.com/raillen/calmtv/internal/web"
)

//go:embed assets/home.ui
var homeUI []byte

//go:embed assets/theme.css
var themeCSS []byte

type App struct {
	window        *gtk.Window
	root          *gtk.Box
	stack         *gtk.Stack
	status        *gtk.Label
	buttons       map[string]*gtk.Button
	screenButtons map[string]map[string]*gtk.Button
	input         *input.Manager
	focus         *focus.Manager
	system        system.Facade
	media         *media.Runtime
	mpris         media.MPRIS
	state         *state.Store
	diagnostics   diagnostics.Reporter
	configRoot    string
	activeScreen  string
	fullscreen    bool
	iptvCatalog   *iptv.Catalog
	appManager    *appmanager.Manager
	registeredApp map[string]bool
	activeAppID   string
	streaming     *exec.Cmd
	activeMedia   string
	volume        int
	muted         bool
}

func New() (*App, error) {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, fmt.Errorf("create GTK builder: %w", err)
	}
	if err = builder.AddFromString(string(homeUI)); err != nil {
		return nil, fmt.Errorf("load Home GtkBuilder: %w", err)
	}

	window, err := object[*gtk.Window](builder, "main_window")
	if err != nil {
		return nil, err
	}
	root, err := object[*gtk.Box](builder, "root_box")
	if err != nil {
		return nil, err
	}
	homeGrid, err := object[*gtk.Grid](builder, "home_grid")
	if err != nil {
		return nil, err
	}
	status, err := object[*gtk.Label](builder, "status_label")
	if err != nil {
		return nil, err
	}

	buttonIDs := []string{"settings", "media", "nanotube", "iptv", "games", "files", "diagnostics", "power", "streaming"}
	buttons := make(map[string]*gtk.Button, len(buttonIDs))
	for _, id := range buttonIDs {
		button, buttonErr := object[*gtk.Button](builder, "card-"+id)
		if buttonErr != nil {
			return nil, buttonErr
		}
		buttons[id] = button
	}

	nodes := make([]focus.Node, 0, len(buttonIDs))
	for index, id := range buttonIDs {
		nodes = append(nodes, focus.Node{ID: id, Row: index / 4, Column: index % 4, Enabled: true})
	}
	focusManager, err := focus.NewManager(nodes, "settings")
	if err != nil {
		return nil, fmt.Errorf("create focus manager: %w", err)
	}

	app := &App{
		window:        window,
		root:          root,
		status:        status,
		buttons:       buttons,
		screenButtons: make(map[string]map[string]*gtk.Button),
		input:         input.NewManager(),
		focus:         focusManager,
		system:        system.NewFacade(system.ExecRunner{}),
		mpris:         media.NewMPRIS(system.ExecRunner{}),
		diagnostics:   diagnostics.NewReporter(diagnostics.ExecCommand{}),
		configRoot:    homeDirectory(filepath.Join(".config", "tv-shell")),
		activeScreen:  "home",
		fullscreen:    os.Getenv("TV_SHELL_WINDOWED") != "1",
		iptvCatalog:   iptv.NewCatalog(),
		appManager:    appmanager.NewManager(appmanager.SystemdStarter{}, nil),
		registeredApp: make(map[string]bool),
		volume:        50,
	}
	dataRoot := homeDirectory(filepath.Join(".local", "share", "tv-shell"))
	if err := os.MkdirAll(dataRoot, 0750); err != nil {
		return nil, fmt.Errorf("create state directory: %w", err)
	}
	app.state, err = state.Open(filepath.Join(dataRoot, "state.db"))
	if err != nil {
		return nil, err
	}
	app.media = media.DefaultRuntime(runtimeSocketPath())
	if err := app.configureScreens(homeGrid); err != nil {
		_ = app.state.Close()
		return nil, err
	}
	app.configure()
	if os.Getenv("TV_SHELL_RECOVERY") == "1" {
		_ = recovery.NewManager(app.configRoot, func(context.Context, string, ...string) error { return nil }).Execute(context.Background(), recovery.ResetUI)
		status.SetText("Modo de recuperação: Home carregada com estado visual mínimo")
	}
	return app, nil
}

func (a *App) Run() {
	a.window.ShowAll()
	if a.fullscreen {
		a.fullscreenOnConfiguredMonitor()
	}
	a.syncFocus()
	gtk.Main()
}

func (a *App) configure() {
	a.window.SetName("calm-tv-window")
	a.window.SetDefaultSize(1280, 720)

	provider, err := gtk.CssProviderNew()
	if err == nil && provider.LoadFromData(string(themeCSS)) == nil {
		if screen, screenErr := gdk.ScreenGetDefault(); screenErr == nil {
			gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
		}
	}

	a.window.Connect("destroy", func() {
		a.saveMediaProgress()
		_ = a.media.Stop()
		a.stopApps()
		_ = a.state.Close()
		gtk.MainQuit()
	})
	a.window.Connect("key-press-event", func(_ *gtk.Window, event *gdk.Event) bool {
		keyEvent := gdk.EventKeyNewFromEvent(event)
		action, ok := input.FromKey(gdk.KeyValName(keyEvent.KeyVal()))
		if !ok {
			return false
		}
		a.input.Emit(input.Event{Action: action, Source: "gtk-keyboard"})
		return true
	})
	a.input.Subscribe(a.handleInput)

	for id, button := range a.buttons {
		buttonID := id
		button.Connect("clicked", func() { a.activate(buttonID) })
	}
	for screen, buttons := range a.screenButtons {
		for id, button := range buttons {
			screenID, buttonID := screen, id
			button.Connect("clicked", func() { a.activateScreenButton(screenID, buttonID) })
		}
	}
}

func (a *App) handleInput(event input.Event) {
	switch event.Action {
	case input.NavUp, input.NavDown, input.NavLeft, input.NavRight:
		a.focus.Move(event.Action)
		a.syncFocus()
	case input.Accept:
		a.activate(a.focus.Current())
	case input.Home:
		a.showHome()
	case input.Menu:
		if a.activeScreen == "home" {
			a.showScreen("diagnostics")
		} else {
			a.status.SetText("Menu avançado disponível na Home")
		}
	case input.Back:
		if a.activeScreen == "home" {
			a.status.SetText("Home — Back")
		} else {
			a.showHome()
		}
	case input.ChannelUp:
		if a.activeScreen == "iptv" {
			a.selectIPTV(1)
		}
	case input.ChannelDown:
		if a.activeScreen == "iptv" {
			a.selectIPTV(-1)
		}
	case input.PlayPause:
		a.togglePlayback()
	case input.VolUp:
		a.adjustVolume(5)
	case input.VolDown:
		a.adjustVolume(-5)
	case input.Mute:
		a.toggleMute()
	case input.Next:
		if err := a.mpris.Next(context.Background()); err != nil {
			a.status.SetText("Próxima faixa indisponível")
		}
	case input.Previous:
		if err := a.mpris.Previous(context.Background()); err != nil {
			a.status.SetText("Faixa anterior indisponível")
		}
	default:
		if number, ok := input.ChannelNumber(event.Action); ok && a.activeScreen == "iptv" {
			a.selectIPTVNumber(number)
			return
		}
		a.status.SetText(string(event.Action))
	}
}

func (a *App) syncFocus() {
	current := a.focus.Current()
	buttons := a.buttons
	if a.activeScreen != "home" {
		buttons = a.screenButtons[a.activeScreen]
	}
	for id, button := range buttons {
		button.SetStateFlags(gtk.STATE_FLAG_NORMAL, true)
		if id == current {
			button.GrabFocus()
		}
	}
	a.status.SetText("Foco: " + current)
}

func (a *App) activate(id string) {
	if id == "" {
		return
	}
	a.showScreen(id)
}

func (a *App) configureScreens(homeGrid *gtk.Grid) error {
	stack, err := gtk.StackNew()
	if err != nil {
		return err
	}
	stack.SetHExpand(true)
	stack.SetVExpand(true)
	a.root.Remove(homeGrid)
	stack.AddNamed(homeGrid, "home")
	a.root.PackStart(stack, true, true, 0)
	a.stack = stack

	screens := map[string][]string{
		"settings":    {"wifi", "wifi-connect", "bluetooth", "bluetooth-power", "volume", "audio-output", "display", "display-apply", "storage", "storage-mount", "suspend", "reboot", "power-off", "back"},
		"media":       {"videos", "music", "usb", "back"},
		"nanotube":    {"continue-watching", "search", "back"},
		"iptv":        {"import-playlist", "favorites", "epg", "back"},
		"games":       {"rom-library", "back"},
		"files":       {"downloads", "videos", "music", "games", "usb", "back"},
		"diagnostics": {"run-diagnostics", "terminal", "back"},
		"power":       {"suspend", "reboot", "power-off", "back"},
		"streaming":   {"open-streaming", "back"},
	}
	labels := map[string]string{
		"wifi": "Wi-Fi", "wifi-connect": "Conectar Wi-Fi padrão", "bluetooth": "Bluetooth", "bluetooth-power": "Ativar Bluetooth", "volume": "Volume", "audio-output": "Saída de áudio", "display": "Display", "display-apply": "Aplicar display padrão", "storage": "Armazenamento USB", "storage-mount": "Montar USB padrão", "suspend": "Suspender", "reboot": "Reiniciar", "power-off": "Desligar", "back": "Voltar", "videos": "Vídeos", "music": "Música", "usb": "USB", "continue-watching": "Continuar assistindo", "search": "Pesquisar", "import-playlist": "Importar playlist", "favorites": "Favoritos", "epg": "Guia EPG", "rom-library": "Biblioteca de jogos", "downloads": "Downloads", "games": "Jogos", "run-diagnostics": "Executar diagnóstico", "terminal": "Terminal avançado", "open-streaming": "Abrir streaming",
	}
	for screen, ids := range screens {
		page, pageErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 18)
		if pageErr != nil {
			return pageErr
		}
		page.SetMarginStart(56)
		page.SetMarginEnd(56)
		page.SetMarginTop(42)
		page.SetMarginBottom(42)
		title, titleErr := gtk.LabelNew(screenTitle(screen))
		if titleErr != nil {
			return titleErr
		}
		title.SetXAlign(0)
		page.PackStart(title, false, false, 0)
		grid, gridErr := gtk.GridNew()
		if gridErr != nil {
			return gridErr
		}
		grid.SetRowSpacing(18)
		grid.SetColumnSpacing(18)
		grid.SetVExpand(true)
		buttons := make(map[string]*gtk.Button, len(ids))
		for index, id := range ids {
			button, buttonErr := gtk.ButtonNewWithLabel(labels[id])
			if buttonErr != nil {
				return buttonErr
			}
			button.SetHExpand(true)
			button.SetVExpand(true)
			button.SetCanFocus(true)
			grid.Attach(button, index%3, index/3, 1, 1)
			buttons[id] = button
		}
		page.PackStart(grid, true, true, 0)
		stack.AddNamed(page, screen)
		a.screenButtons[screen] = buttons
	}
	return nil
}

func screenTitle(screen string) string {
	return map[string]string{"settings": "Quick Settings", "media": "Mídia local", "nanotube": "NanoTube", "iptv": "IPTV", "games": "Jogos", "files": "Arquivos e downloads", "diagnostics": "Diagnóstico", "power": "Energia", "streaming": "Streaming opcional"}[screen]
}

func (a *App) showScreen(screen string) {
	buttons, ok := a.screenButtons[screen]
	if !ok || len(buttons) == 0 {
		return
	}
	nodes := make([]focus.Node, 0, len(buttons))
	for index, id := range orderedButtonIDs(buttons) {
		nodes = append(nodes, focus.Node{ID: id, Row: index / 3, Column: index % 3, Enabled: true})
	}
	if err := a.focus.ReplaceScreen(nodes, nodes[0].ID); err != nil {
		a.status.SetText("Erro de foco")
		return
	}
	a.activeScreen = screen
	a.stack.SetVisibleChildName(screen)
	a.status.SetText("Tela: " + screen)
	a.syncFocus()
}

func (a *App) showHome() {
	a.saveMediaProgress()
	_ = a.media.Stop()
	a.stopApps()
	ids := []string{"settings", "media", "nanotube", "iptv", "games", "files", "diagnostics", "power", "streaming"}
	nodes := make([]focus.Node, 0, len(ids))
	for index, id := range ids {
		nodes = append(nodes, focus.Node{ID: id, Row: index / 4, Column: index % 4, Enabled: true})
	}
	_ = a.focus.ReplaceScreen(nodes, "settings")
	a.activeScreen = "home"
	a.stack.SetVisibleChildName("home")
	a.status.SetText("Home")
	a.syncFocus()
}

func (a *App) activateScreenButton(screen, id string) {
	if id == "back" {
		a.showHome()
		return
	}
	if (screen == "settings" && (id == "wifi" || id == "bluetooth" || id == "volume" || id == "audio-output" || id == "display" || id == "storage" || id == "suspend" || id == "reboot" || id == "power-off")) || screen == "power" || id == "usb" {
		result, err := a.system.Execute(context.Background(), id)
		if err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText(result.Summary)
		return
	}
	if screen == "settings" && id == "wifi-connect" {
		networks, err := a.system.Network.List(context.Background())
		if err != nil || len(networks) == 0 {
			a.status.SetText("Nenhuma rede Wi-Fi disponível")
			return
		}
		name := os.Getenv("TV_SHELL_WIFI_CONNECTION")
		if name == "" {
			name = networks[0].Name
		}
		if err := a.system.Network.Connect(context.Background(), name); err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("Wi-Fi conectado: " + name)
		return
	}
	if screen == "settings" && id == "bluetooth-power" {
		if err := a.system.Bluetooth.SetPowered(context.Background(), true); err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("Bluetooth ativado")
		return
	}
	if screen == "settings" && id == "display-apply" {
		modes, err := a.system.Display.ListModes(context.Background())
		if err != nil || len(modes) == 0 {
			a.status.SetText("Nenhum modo de display disponível")
			return
		}
		mode := modes[0]
		if output := os.Getenv("TV_SHELL_DISPLAY_OUTPUT"); output != "" {
			mode.Output = output
		}
		if selectedMode := os.Getenv("TV_SHELL_DISPLAY_MODE"); selectedMode != "" {
			mode.Mode = selectedMode
		}
		if err := a.system.Display.SetMode(context.Background(), mode.Output, mode.Mode); err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("Display aplicado: " + mode.Output + " " + mode.Mode)
		return
	}
	if screen == "settings" && id == "storage-mount" {
		volumes, err := a.system.Storage.ListRemovable(context.Background())
		if err != nil || len(volumes) == 0 {
			a.status.SetText("Nenhum USB disponível")
			return
		}
		if err := a.system.Storage.Mount(context.Background(), volumes[0].Device); err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("USB montado: " + volumes[0].Device)
		return
	}
	if screen == "media" && (id == "videos" || id == "music") {
		root := homeDirectory(map[string]string{"videos": "Videos", "music": "Music"}[id])
		items, err := library.Scan(root)
		if err != nil || len(items) == 0 {
			a.status.SetText("Nenhuma mídia encontrada em " + root)
			return
		}
		for _, item := range items {
			_ = a.state.SaveMediaItem(context.Background(), state.MediaItem{ID: item.Path, Path: item.Path, Title: item.Title, Kind: item.Kind})
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = a.media.Open(ctx, items[0].Path)
		cancel()
		if err != nil {
			a.status.SetText("Não foi possível reproduzir: " + err.Error())
			return
		}
		a.activeMedia = items[0].Path
		if position, _, progressErr := a.state.RestoreProgress(context.Background(), items[0].Path); progressErr == nil && position > 0 {
			_ = a.media.Client().Seek(context.Background(), position)
		}
		a.status.SetText("Reproduzindo: " + items[0].Title)
		return
	}
	if screen == "files" && (id == "downloads" || id == "videos" || id == "music" || id == "games") {
		folder := map[string]string{"downloads": "Downloads", "videos": "Videos", "music": "Music", "games": "Games"}[id]
		entries, err := os.ReadDir(homeDirectory(folder))
		if err != nil {
			a.status.SetText(folder + ": diretório indisponível")
			return
		}
		a.status.SetText(fmt.Sprintf("%s: %d itens", folder, len(entries)))
		return
	}
	if screen == "nanotube" && (id == "continue-watching" || id == "search") {
		if a.appManager.Running("nanotube") {
			a.status.SetText("NanoTube já está aberto")
			return
		}
		binary := os.Getenv("NANOTUBE_BINARY")
		if binary == "" {
			binary = "nanotube"
		}
		if err := a.launchManaged(context.Background(), "nanotube", appmanager.Manifest{ID: "nanotube", Command: binary, ResourceClass: appmanager.Heavy}); err != nil {
			a.status.SetText("NanoTube indisponível: " + err.Error())
			return
		}
		a.status.SetText("NanoTube aberto sob demanda")
		return
	}
	if screen == "streaming" && id == "open-streaming" {
		if a.streaming != nil {
			a.status.SetText("Streaming já está aberto")
			return
		}
		serviceURL := os.Getenv("TV_SHELL_STREAMING_URL")
		if serviceURL == "" {
			a.status.SetText("Configure TV_SHELL_STREAMING_URL para abrir um serviço")
			return
		}
		launcher := web.NewLauncher(os.Getenv("FIREFOX_ESR_BINARY"), filepath.Join(a.configRoot, "browser"))
		command, err := launcher.Start(context.Background(), serviceURL)
		if err != nil {
			a.status.SetText("Firefox ESR indisponível: " + err.Error())
			return
		}
		a.streaming = command
		a.status.SetText("Streaming aberto em kiosk")
		return
	}
	if screen == "iptv" && id == "import-playlist" {
		playlist := filepath.Join(a.configRoot, "iptv", "playlist.m3u")
		file, err := os.Open(playlist)
		if err != nil {
			a.status.SetText("Playlist IPTV não configurada")
			return
		}
		channels, parseErr := iptv.ParseM3U(file)
		_ = file.Close()
		if parseErr != nil {
			a.status.SetText("Playlist IPTV inválida")
			return
		}
		a.iptvCatalog.Replace(channels)
		for _, channel := range channels {
			_ = a.state.SaveChannel(context.Background(), channel.ID, channel.Name, channel.Group, channel.Logo, channel.StreamURL)
		}
		a.status.SetText(fmt.Sprintf("IPTV: %d canais importados", len(channels)))
		if len(channels) > 0 {
			a.playIPTVChannel(channels[0])
		}
		return
	}
	if screen == "iptv" && id == "favorites" {
		channel, err := a.iptvCatalog.ToggleFavorite()
		if err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("Favorito: " + channel.Name)
		_ = a.state.SaveFavorite(context.Background(), channel.ID+"\x00"+channel.StreamURL, "channel", a.iptvCatalog.IsFavorite(channel))
		return
	}
	if screen == "iptv" && id == "epg" {
		count, err := a.readEPG()
		if err != nil {
			a.status.SetText("EPG indisponível: " + err.Error())
			return
		}
		a.status.SetText(fmt.Sprintf("EPG: %d programas processados", count))
		return
	}
	if screen == "games" && id == "rom-library" {
		a.launchFirstROM()
		return
	}
	if screen == "diagnostics" && id == "run-diagnostics" {
		report := a.diagnostics.Collect(context.Background())
		reportPath := homeDirectory(filepath.Join(".local", "share", "tv-shell", "diagnostics", "latest.json"))
		if err := diagnostics.Save(reportPath, report); err != nil {
			a.status.SetText("Diagnóstico concluído, mas não foi salvo: " + err.Error())
			return
		}
		a.status.SetText(fmt.Sprintf("Diagnóstico salvo: %d verificações", len(report.Commands)))
		return
	}
	if screen == "diagnostics" && id == "terminal" {
		if err := launchMaintenanceTerminal(); err != nil {
			a.status.SetText(err.Error())
			return
		}
		a.status.SetText("Terminal avançado aberto como usuário normal")
		return
	}
	a.status.SetText(screen + ": " + id)
}

func (a *App) stopApps() {
	if a.activeAppID != "" {
		_ = a.appManager.Stop(context.Background(), a.activeAppID)
		a.activeAppID = ""
	}
	if a.streaming != nil {
		_ = a.streaming.Process.Kill()
		a.streaming = nil
	}
}

func (a *App) launchManaged(ctx context.Context, appID string, manifest appmanager.Manifest) error {
	if !a.registeredApp[appID] {
		if err := a.appManager.Register(manifest); err != nil {
			return err
		}
		a.registeredApp[appID] = true
	}
	if err := a.appManager.Launch(ctx, appID); err != nil {
		return err
	}
	a.activeAppID = appID
	return nil
}

func (a *App) saveMediaProgress() {
	if a.activeMedia == "" || a.media.Client() == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	position, err := a.media.Client().Position(ctx)
	if err != nil {
		return
	}
	duration, err := a.media.Client().Duration(ctx)
	if err != nil {
		return
	}
	_ = a.state.SaveProgress(ctx, a.activeMedia, position, duration)
	a.activeMedia = ""
}

func (a *App) togglePlayback() {
	client := a.media.Client()
	if client == nil {
		a.status.SetText("Nenhuma mídia em reprodução")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var err error
	if client.State() == media.StatePlaying {
		err = client.Pause(ctx)
	} else {
		err = client.Play(ctx)
	}
	if err != nil {
		a.status.SetText("Controle de mídia indisponível: " + err.Error())
		return
	}
	a.status.SetText("Mídia: " + string(client.State()))
}

func (a *App) adjustVolume(delta int) {
	a.volume += delta
	if a.volume < 0 {
		a.volume = 0
	}
	if a.volume > 100 {
		a.volume = 100
	}
	if err := a.system.Audio.SetVolume(context.Background(), "@DEFAULT_AUDIO_SINK@", a.volume); err != nil {
		a.status.SetText(err.Error())
		return
	}
	a.muted = false
	a.status.SetText(fmt.Sprintf("Volume: %d%%", a.volume))
}

func (a *App) toggleMute() {
	a.muted = !a.muted
	if err := a.system.Audio.SetMute(context.Background(), "@DEFAULT_AUDIO_SINK@", a.muted); err != nil {
		a.status.SetText(err.Error())
		return
	}
	if a.muted {
		a.status.SetText("Áudio silenciado")
	} else {
		a.status.SetText("Áudio restaurado")
	}
}

func (a *App) selectIPTV(delta int) {
	channel, err := a.iptvCatalog.Move(delta)
	if err != nil {
		a.status.SetText(err.Error())
		return
	}
	a.playIPTVChannel(channel)
}

func (a *App) selectIPTVNumber(number int) {
	channel, err := a.iptvCatalog.SelectNumber(number)
	if err != nil {
		a.status.SetText(err.Error())
		return
	}
	a.playIPTVChannel(channel)
}

func (a *App) playIPTVChannel(channel iptv.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := a.media.OpenURL(ctx, channel.StreamURL)
	cancel()
	if err != nil {
		a.status.SetText("Canal selecionado: " + channel.Name)
		return
	}
	a.status.SetText("Canal: " + channel.Name)
}

func (a *App) readEPG() (int, error) {
	file, err := os.Open(filepath.Join(a.configRoot, "iptv", "guide.xml"))
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	err = iptv.ParseXMLTV(file, nil, func(program iptv.Program) error {
		count++
		return a.state.SaveEPGProgram(context.Background(), program.ChannelID, program.Start, program.End, program.Title, program.Description)
	})
	return count, err
}

func (a *App) launchFirstROM() {
	root := homeDirectory("Games")
	entries, err := os.ReadDir(root)
	if err != nil {
		a.status.SetText("Biblioteca de jogos indisponível")
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(root, entry.Name())
		file, openErr := os.Open(path)
		if openErr != nil {
			continue
		}
		rom, classifyErr := games.Classify(path, file)
		_ = file.Close()
		if classifyErr != nil || rom.Core == "" {
			continue
		}
		binary := os.Getenv("RETROARCH_BINARY")
		if binary == "" {
			binary = "retroarch"
		}
		appID := "retroarch-" + rom.Hash[:12]
		manifest := appmanager.Manifest{ID: appID, Command: binary, Args: []string{"-L", rom.Core, rom.Path, "--save-dir", filepath.Join(root, ".saves")}, ResourceClass: appmanager.ExclusiveHeavy, Exclusive: true}
		if err := a.launchManaged(context.Background(), appID, manifest); err != nil {
			a.status.SetText("RetroArch indisponível: " + err.Error())
			return
		}
		a.status.SetText("Jogo iniciado: " + rom.Title)
		return
	}
	a.status.SetText("Nenhuma ROM compatível em " + root)
}

func homeDirectory(child string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return child
	}
	return filepath.Join(home, child)
}

func runtimeSocketPath() string {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		runtimeDir = filepath.Join(os.TempDir(), "tv-shell-"+strconv.Itoa(os.Getuid()))
	}
	return filepath.Join(runtimeDir, "tv-shell-mpv.sock")
}

func orderedButtonIDs(buttons map[string]*gtk.Button) []string {
	ids := make([]string, 0, len(buttons))
	for _, id := range []string{"wifi", "wifi-connect", "bluetooth", "bluetooth-power", "volume", "audio-output", "display", "display-apply", "storage", "storage-mount", "suspend", "reboot", "power-off", "videos", "music", "usb", "continue-watching", "search", "import-playlist", "favorites", "epg", "rom-library", "downloads", "games", "run-diagnostics", "terminal", "open-streaming", "back"} {
		if _, ok := buttons[id]; ok {
			ids = append(ids, id)
		}
	}
	return ids
}

func launchMaintenanceTerminal() error {
	for _, candidate := range []string{"qterminal", "xfce4-terminal", "mate-terminal", "konsole", "xterm"} {
		path, err := exec.LookPath(candidate)
		if err != nil {
			continue
		}
		command := exec.Command(path)
		if err := command.Start(); err != nil {
			return fmt.Errorf("não foi possível abrir o terminal avançado: %w", err)
		}
		go func() { _ = command.Wait() }()
		return nil
	}
	return fmt.Errorf("nenhum terminal gráfico encontrado; instale qterminal ou xterm")
}

func object[T any](builder *gtk.Builder, id string) (T, error) {
	var zero T
	value, err := builder.GetObject(id)
	if err != nil {
		return zero, fmt.Errorf("get GTK object %q: %w", id, err)
	}
	result, ok := value.(T)
	if !ok {
		return zero, fmt.Errorf("GTK object %q has unexpected type", id)
	}
	return result, nil
}
