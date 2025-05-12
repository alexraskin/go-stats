package plugins

// all plugins must implement this interface
type Plugin interface {
	Name() string
	Metrics() (map[string]any, error)
}
