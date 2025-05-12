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
