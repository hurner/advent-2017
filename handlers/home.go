package handlers

import (
	"time"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// home returns a simple HTTP handler function which writes a response.
func home(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {

		hostname := os.Getenv("HOSTNAME")
		time := time.Now().Format("2006-01-02 15:04:05")

		info := struct {
			BuildTime string `json:"buildTime"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
			Hostname  string `json:"hostname"`
			Time      string `json:"time"`
		}{
			buildTime, commit, release, hostname, time,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
