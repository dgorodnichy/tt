package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func address() string {
	addr := ":8080"

	if val := os.Getenv("PORT"); val != "" {
		addr = fmt.Sprintf(":%s", val)
	}
	return addr
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		status, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			log.Printf("incorrect json: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if _, err = w.Write(status); err != nil {
			log.Printf("network problems: %v", err)
		}
	})

	log.Fatal(http.ListenAndServe(address(), mux))
}
