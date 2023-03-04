package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "80", "")
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		data := struct {
			Hostname string      `json:"hostname,omitempty"`
			IP       []string    `json:"ip,omitempty"`
			Headers  http.Header `json:"header,omitempty"`
			URL      string      `json:"url,omitempty"`
			Host     string      `json:"host,omitempty"`
			Method   string      `json:"method,omitempty"`
		}{
			Hostname: hostname,
			IP:       []string{},
			Headers:  r.Header,
			URL:      r.URL.RequestURI(),
			Host:     r.Host,
			Method:   r.Method,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	mux.Handle("/sin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mux.Handle("/iad", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mux.Handle("/cdg", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
