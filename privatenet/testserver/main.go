package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

	mux.Handle("/sin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://testserver.sin.internal/api")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, string(body))
		if err := r.Write(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	mux.Handle("/iad", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://testserver.iad.internal/api")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, string(body))
		if err := r.Write(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	mux.Handle("/cdg", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://testserver.cdg.internal/api")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, string(body))
		if err := r.Write(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
