package appmanager

import (
	"context"
	"fmt"
	"sync"
)

type Manifest struct {
	ID                string
	Command           string
	Args              []string
	ResourceClass     ResourceClass
	BackgroundAllowed bool
	Exclusive         bool
}

type State struct {
	Route       string `json:"route"`
	SelectedID  string `json:"selected_id"`
	PositionSec int64  `json:"position_sec"`
}

type StateStore interface {
	Save(ctx context.Context, appID string, state State) error
	Restore(ctx context.Context, appID string) (State, error)
}

type Process interface {
	Wait() error
	Stop(ctx context.Context) error
}

type Starter interface {
	Start(ctx context.Context, manifest Manifest, policy ResourcePolicy) (Process, error)
}

type Manager struct {
	mu         sync.Mutex
	manifests  map[string]Manifest
	policies   map[string]ResourcePolicy
	processes  map[string]Process
	states     StateStore
	starter    Starter
	foreground string
}

func NewManager(starter Starter, states StateStore) *Manager {
	return &Manager{
		manifests: make(map[string]Manifest),
		policies:  make(map[string]ResourcePolicy),
		processes: make(map[string]Process),
		states:    states,
		starter:   starter,
	}
}

func (m *Manager) Register(manifest Manifest) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if manifest.ID == "" || manifest.Command == "" {
		return fmt.Errorf("app manifest requires id and command")
	}
	if _, exists := m.manifests[manifest.ID]; exists {
		return fmt.Errorf("app %q is already registered", manifest.ID)
	}
	if manifest.ResourceClass == "" {
		manifest.ResourceClass = Medium
	}
	m.manifests[manifest.ID] = manifest
	return nil
}

// Background records that an app may remain available as a background
// process. Callers must opt in per manifest; persisted app state otherwise
// remains the preferred representation of recent apps.
func (m *Manager) Background(appID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	manifest, ok := m.manifests[appID]
	if !ok {
		return fmt.Errorf("app %q is not registered", appID)
	}
	if !manifest.BackgroundAllowed {
		return fmt.Errorf("app %q does not allow background execution", appID)
	}
	if _, running := m.processes[appID]; !running {
		return fmt.Errorf("app %q is not running", appID)
	}
	if m.foreground == appID {
		m.foreground = ""
	}
	return nil
}

func (m *Manager) ResourcePolicy(appID string) (ResourcePolicy, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	manifest, ok := m.manifests[appID]
	if !ok {
		return ResourcePolicy{}, fmt.Errorf("app %q is not registered", appID)
	}
	if policy, ok := m.policies[appID]; ok {
		return policy, nil
	}
	return DefaultPolicy(manifest.ResourceClass), nil
}

func (m *Manager) SetResourcePolicy(appID string, policy ResourcePolicy) error {
	if err := policy.Validate(); err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.manifests[appID]; !ok {
		return fmt.Errorf("app %q is not registered", appID)
	}
	m.policies[appID] = policy
	return nil
}

func (m *Manager) Launch(ctx context.Context, appID string) error {
	m.mu.Lock()
	manifest, ok := m.manifests[appID]
	if !ok {
		m.mu.Unlock()
		return fmt.Errorf("app %q is not registered", appID)
	}
	if _, running := m.processes[appID]; running {
		m.foreground = appID
		m.mu.Unlock()
		return nil
	}
	if manifest.Exclusive || manifest.ResourceClass == ExclusiveHeavy {
		for runningID := range m.processes {
			if runningID != appID {
				m.mu.Unlock()
				if err := m.Stop(ctx, runningID); err != nil {
					return err
				}
				m.mu.Lock()
			}
		}
	}
	policy := DefaultPolicy(manifest.ResourceClass)
	if configured, configuredOK := m.policies[appID]; configuredOK {
		policy = configured
	}
	process, err := m.starter.Start(ctx, manifest, policy)
	if err != nil {
		m.mu.Unlock()
		return fmt.Errorf("launch %q: %w", appID, err)
	}
	m.processes[appID] = process
	m.foreground = appID
	m.mu.Unlock()
	return nil
}

func (m *Manager) Stop(ctx context.Context, appID string) error {
	m.mu.Lock()
	process, ok := m.processes[appID]
	if !ok {
		m.mu.Unlock()
		return nil
	}
	delete(m.processes, appID)
	if m.foreground == appID {
		m.foreground = ""
	}
	m.mu.Unlock()
	if err := process.Stop(ctx); err != nil {
		return fmt.Errorf("stop %q: %w", appID, err)
	}
	return nil
}

func (m *Manager) Foreground() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.foreground
}

func (m *Manager) SaveState(ctx context.Context, appID string, state State) error {
	if m.states == nil {
		return nil
	}
	return m.states.Save(ctx, appID, state)
}

func (m *Manager) Restore(ctx context.Context, appID string) (State, error) {
	if m.states == nil {
		return State{}, nil
	}
	return m.states.Restore(ctx, appID)
}

func (m *Manager) Running(appID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.processes[appID]
	return ok
}
