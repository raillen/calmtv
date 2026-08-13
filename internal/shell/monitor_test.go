package shell

import "testing"

func TestResolveMonitorIndexDefaultsToPrimary(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		count    int
		primary  int
		expected int
	}{
		{name: "empty", value: "", count: 2, primary: 1, expected: 1},
		{name: "primary", value: "primary", count: 2, primary: 1, expected: 1},
		{name: "explicit monitor", value: "0", count: 2, primary: 1, expected: 0},
		{name: "invalid value", value: "external", count: 2, primary: 1, expected: 1},
		{name: "out of range", value: "2", count: 2, primary: 1, expected: 1},
		{name: "invalid primary", value: "", count: 2, primary: 9, expected: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := resolveMonitorIndex(test.value, test.count, test.primary); got != test.expected {
				t.Fatalf("resolveMonitorIndex(%q, %d, %d) = %d, want %d", test.value, test.count, test.primary, got, test.expected)
			}
		})
	}
}
