package main

import (
	"flag"
	"log"
	"net/http"
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
	mux.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mux.Handle("/sin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mux.Handle("/iad", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mux.Handle("/cdg", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
