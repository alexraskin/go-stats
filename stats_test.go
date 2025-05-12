package stats

import (
	"testing"
)

type mockPlugin struct{}

func (m *mockPlugin) Name() string {
	return "mock"
}

func (m *mockPlugin) Metrics() (map[string]any, error) {
	return map[string]any{
		"mock_metric": 42,
	}, nil
}

func TestNewStatsWithPlugin(t *testing.T) {
	mock := &mockPlugin{}
	stats, err := NewStats(mock)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stats["mock_metric"] != 42 {
		t.Fatalf("expected mock_metric to be 42, got %v", stats["mock_metric"])
	}
}
