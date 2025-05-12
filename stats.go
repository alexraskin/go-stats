package stats

import (
	"maps"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
)

var startTime = time.Now()

type Plugin interface {
	Name() string
	Metrics() (map[string]any, error)
}

func NewStats(plugins ...Plugin) (map[string]any, error) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	data := map[string]any{
		"go":                runtime.Version(),
		"uptime":            formatDuration(time.Since(startTime)),
		"memory_used":       humanize.Bytes(stats.Alloc),
		"total_memory":      humanize.Bytes(stats.Sys),
		"garbage_collected": humanize.Bytes(stats.TotalAlloc),
		"num_goroutines":    runtime.NumGoroutine(),
	}

	for _, plugin := range plugins {
		metrics, err := plugin.Metrics()
		if err != nil {
			return nil, err
		}
		maps.Copy(data, metrics)
	}

	return data, nil
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return humanize.Comma(int64(hours)) + "h " +
		humanize.Comma(int64(minutes)) + "m " +
		humanize.Comma(int64(seconds)) + "s"
}
