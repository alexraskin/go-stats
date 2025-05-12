# go-stats

Simple Go package for collecting runtime statistics with optional plugins.

## Installation

```bash
go get github.com/alexraskin/go-stats
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexraskin/go-stats"
)

func main() {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		stats, err := stats.NewStats()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(stats)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
```

## Plugins

The `go-stats` package supports optional plugins. To create a plugin, implement the `stats.Plugin` interface:

```go
package myplugin

import "github.com/alexraskin/go-stats"

type MyPlugin struct{}

func (p *MyPlugin) Name() string {
	return "myplugin"
}

func (p *MyPlugin) Metrics() (map[string]any, error) {
	return map[string]any{
		"metric1": 100,
		"metric2": 200,
	}, nil
}
```
