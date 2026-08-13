package focus

import (
	"sort"

	"github.com/raillen/calmtv/internal/input"
)

// Node is a logical focus target. Screens describe targets; FocusManager owns
// movement, disabled-item skipping and the current target.
type Node struct {
	ID      string
	Row     int
	Column  int
	Enabled bool
}

type Manager struct {
	nodes   map[string]Node
	current string
}

func NewManager(nodes []Node, initial string) (*Manager, error) {
	manager := &Manager{nodes: make(map[string]Node)}
	for _, node := range nodes {
		if node.ID == "" {
			continue
		}
		manager.nodes[node.ID] = node
	}
	if !manager.hasEnabled(initial) {
		initial = manager.firstEnabled()
	}
	if initial == "" {
		return nil, ErrNoFocusableNode
	}
	manager.current = initial
	return manager, nil
}

var ErrNoFocusableNode = &focusError{"no enabled focus target exists"}

type focusError struct{ message string }

func (e *focusError) Error() string { return e.message }

func (m *Manager) Current() string { return m.current }

// ReplaceScreen installs the active screen's targets while keeping one
// central navigation owner for the entire Shell.
func (m *Manager) ReplaceScreen(nodes []Node, initial string) error {
	updated := make(map[string]Node, len(nodes))
	for _, node := range nodes {
		if node.ID != "" {
			updated[node.ID] = node
		}
	}
	m.nodes = updated
	if !m.hasEnabled(initial) {
		initial = m.firstEnabled()
	}
	if initial == "" {
		return ErrNoFocusableNode
	}
	m.current = initial
	return nil
}

func (m *Manager) SetEnabled(id string, enabled bool) bool {
	node, ok := m.nodes[id]
	if !ok {
		return false
	}
	node.Enabled = enabled
	m.nodes[id] = node
	if !enabled && m.current == id {
		m.current = m.firstEnabled()
	}
	return true
}

func (m *Manager) Move(action input.Action) string {
	dRow, dColumn, ok := direction(action)
	if !ok {
		return m.current
	}
	current, ok := m.nodes[m.current]
	if !ok {
		m.current = m.firstEnabled()
		return m.current
	}

	candidates := make([]Node, 0, len(m.nodes))
	for _, node := range m.nodes {
		if !node.Enabled || node.ID == current.ID {
			continue
		}
		deltaRow := node.Row - current.Row
		deltaColumn := node.Column - current.Column
		if dRow != 0 && deltaRow*dRow <= 0 {
			continue
		}
		if dColumn != 0 && deltaColumn*dColumn <= 0 {
			continue
		}
		if dRow != 0 && deltaColumn != 0 {
			continue
		}
		if dColumn != 0 && deltaRow != 0 {
			continue
		}
		candidates = append(candidates, node)
	}

	sort.Slice(candidates, func(i, j int) bool {
		iDistance := distance(current, candidates[i])
		jDistance := distance(current, candidates[j])
		if iDistance == jDistance {
			return candidates[i].ID < candidates[j].ID
		}
		return iDistance < jDistance
	})
	if len(candidates) > 0 {
		m.current = candidates[0].ID
	}
	return m.current
}

func direction(action input.Action) (row, column int, ok bool) {
	switch action {
	case input.NavUp:
		return -1, 0, true
	case input.NavDown:
		return 1, 0, true
	case input.NavLeft:
		return 0, -1, true
	case input.NavRight:
		return 0, 1, true
	default:
		return 0, 0, false
	}
}

func distance(a, b Node) int {
	row := a.Row - b.Row
	column := a.Column - b.Column
	if row < 0 {
		row = -row
	}
	if column < 0 {
		column = -column
	}
	return row + column
}

func (m *Manager) hasEnabled(id string) bool {
	node, ok := m.nodes[id]
	return ok && node.Enabled
}

func (m *Manager) firstEnabled() string {
	ids := make([]string, 0, len(m.nodes))
	for id, node := range m.nodes {
		if node.Enabled {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	if len(ids) == 0 {
		return ""
	}
	return ids[0]
}
